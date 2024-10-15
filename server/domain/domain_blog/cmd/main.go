package main

import (
	"domain_blog/common"
	"domain_blog/conf"
	"domain_blog/domain"
	"domain_blog/domain/core"
	"domain_blog/infrastructure/database"
	"domain_blog/infrastructure/grpc/interceptor"
	pb "domain_blog/infrastructure/grpc/pb"
	"domain_blog/internal/cacheUtil"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"
	"net"
)

var (
	configPath = flag.String("config", "domain_config.yaml", "config name")
)

func main() {
	flag.Parse()

	defer func() {
		if err := recover(); err != nil {
			if common.Log != nil {
				common.Log.ErrorMsg("%v", err)
			} else {
				fmt.Printf("panic:%v\n", err)
			}
		}
		Stop()
	}()

	err := Start()
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Cfg.Server.Port))
	if err != nil {
		common.Log.ErrorMsg("failed to listen: %v", err)
		panic(err)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.TranslateCtxInterceptor,
			interceptor.LoggerInterceptor(common.Log),
		),
	)
	pb.RegisterBlogGRPCServer(s, domain.NewBlogGrpcServer())

	var g group.Group
	g.Add(func() error {
		if err := s.Serve(listener); err != nil {
			return fmt.Errorf("%s 服务启动失败: %v", conf.Section, err)
		}
		return nil
	}, func(error) {
		listener.Close()
	})
	common.Log.InfoMsg("%s 服务启动成功 %v", conf.Section, conf.Cfg.Server.Port)
	fmt.Printf("%s 服务启动成功 %v \n", conf.Section, conf.Cfg.Server.Port)
	initCancelInterrupt(&g)
	if err = g.Run(); err != nil {
		common.Log.ErrorMsg("g.Run() exit: %s", err.Error())
		panic(err)
	}
}

func InitLog() {
	var log = common.LoggerConfig{
		Filename:   conf.Cfg.Log.LogPath + "/" + conf.Cfg.Log.Filename,
		LogLevel:   conf.Cfg.Log.Level,
		MaxSize:    conf.Cfg.Log.Maxsize,
		MaxAge:     conf.Cfg.Log.MaxDays,
		MaxBackups: 2,
	}
	common.InitLogger(log)
	common.Log.InfoMsg("日志组件初始化成功")
}

func Start() error {
	cfg, err := conf.Parse(*configPath)
	if err != nil {
		fmt.Printf("解析配置文件错误:%s", err.Error())
		return err
	}
	conf.Cfg = cfg
	InitLog()
	err = database.Init(conf.Cfg.Database)
	if err != nil {
		fmt.Printf("初始化数据库组件错误:%s", err.Error())
		return err
	}
	// 启动临时缓存
	cacheUtil.InitCache()
	// 是否强行初始化admin用户
	if conf.Cfg.System.InitUser == 1 {
		if err := core.NewUserService().ForceCreateAdmin(); err != nil {
			fmt.Printf("创建admin用户失败 err:%s\n", err.Error())
			return err
		}
	} else {
		// admin用户不存在时自动创建
		if !core.NewUserService().IsExistAdmin() {
			if err := core.NewUserService().ForceCreateAdmin(); err != nil {
				fmt.Printf("创建admin用户失败 err:%s\n", err.Error())
				return err
			}
		}
	}
	// 启动redis
	if conf.Cfg.Redis.Enable == 1 {
		if err := RegisterRedis(); err != nil {
			fmt.Printf("连接redis失败 err:%s\n", err.Error())
			return err
		}
	}

	return nil
}

func RegisterRedis() error {
	var addresses []string
	for _, v := range conf.Cfg.Redis.Hosts {
		addresses = append(addresses, fmt.Sprintf("%s:%d", v.Host, v.Port))
	}
	serverInfo := common.RedisConfig{
		Address:  addresses,
		Password: conf.Cfg.Redis.Password,
	}
	if conf.Cfg.Redis.HaType == 0 {
		serverInfo.Type = common.RedisModeOfSignal
	} else {
		serverInfo.Type = common.RedisModeOfCluster
	}
	err := common.InitRedis(serverInfo)
	if err != nil {
		return err
	}
	return nil
}

func Stop() {

	// todo
}

func initCancelInterrupt(g *group.Group) {
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		sig := <-c
		return fmt.Errorf("received signal %s", sig)

	}, func(error) {})
}
