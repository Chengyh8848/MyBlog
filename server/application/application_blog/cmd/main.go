package main

import (
	"application_blog/common"
	"application_blog/conf"
	"application_blog/controller/v1"
	client "application_blog/internal/pkg/grpc"
	"application_blog/internal/utils/cacheUtil"
	"application_blog/middleware"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	configPath = flag.String("config", "application_config.yaml", "config name")
)

func main() {
	flag.Parse()

	defer func() {
		if err := recover(); err != nil {
			if common.Log != nil {
				common.Log.ErrorMsg("%v", err)
			}
		}
		Stop()
	}()

	flag.Parse()

	err := Start()
	if err != nil {
		panic(err)
	}

	engine := gin.New()
	engine.Use(middleware.ThrowPanic())
	//engine.Use(logger.WriteLoggerToFile())
	v1.NoNeedLoginRoute(engine.Group("api"))
	v1.OnlyNeedLoginRoute(engine.Group("api", middleware.AuthToken()))

	common.Log.InfoMsg("服务启动：%v", conf.Cfg.Server.Port)

	err = engine.Run(fmt.Sprintf(":%v", conf.Cfg.Server.Port))
	if err != nil {
		common.Log.ErrorMsg("服务启动失败：%v", err.Error())
		panic(err)
	}
}

func Start() error {
	// 读取配置文件
	cfg, err := conf.Parse(*configPath)
	if err != nil {
		fmt.Printf("解析配置文件错误:%s", err.Error())
		return err
	}
	// 初始化日志
	conf.Cfg = cfg
	InitLog()

	// 启动临时缓存
	cacheUtil.InitCache()

	// GRPC初始化
	grpcConn, err := grpc.Dial(fmt.Sprintf("%s:%d", conf.Cfg.Service.GrpcIp, conf.Cfg.Service.GrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(common.UnaryClientInterceptor),
	)

	if err != nil {
		return err
	}
	client.InitGrpcClients(grpcConn)
	return nil
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

func Stop() {

	// todo
}
