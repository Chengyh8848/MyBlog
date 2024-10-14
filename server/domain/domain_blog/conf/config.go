package conf

import (
	"bytes"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// 0-单机模式 1-哨兵模式 2-集群模式
const (
	RedisModeOfSignal = iota
	RedisModeOfSentinel
	RedisModeOfCluster
)

const (
	ContextReqUUid = "uuid"
	ContextMethod  = "method"
)

const (
	Section = "domain_video"
	UnitM   = 1024 * 1024 //单位M
)

var Cfg *Config

type Config struct {
	Database Database `yaml:"database"` // 数据库配置
	Server   Server   `yaml:"server"`   // 服务器配置
	System   System   `yaml:"system"`   // 系统配置
	Redis    Redis    `yaml:"redis"`    // Redis配置
	Log      Log      `yaml:"logger"`   // 日志配置
	DaServer DaServer `yaml:"daServer"` // DA配置
}

type System struct {
	Threshold string `yaml:"threshold"`
	InitUser  int    `yaml:"initUser"`
}

type Redis struct {
	Enable   int    `yaml:"enable"`
	HaType   int    `yaml:"haType"`
	Seconds  int    `yaml:"seconds"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Host1    string `yaml:"host1"`
	Port1    int    `yaml:"port1"`
	Host2    string `yaml:"host2"`
	Port2    int    `yaml:"port2"`
	Hosts    []HostInfo
}

type HostInfo struct {
	Host string
	Port int
}

type Database struct {
	DbName      string `yaml:"dbName"`
	AutoMigrate bool   `yaml:"autoMigrate"`
}

type Server struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}

type Log struct {
	Typ       string `yaml:"type"`
	Filename  string `yaml:"filename"`
	MaxLines  int    `yaml:"maxlines"`
	Maxsize   int    `yaml:"maxsize"`
	Daily     bool   `yaml:"daily"`
	MaxDays   int    `yaml:"maxdays"`
	Rotate    bool   `yaml:"rotate"`
	Level     string `yaml:"level"`
	LogPath   string `yaml:"logpath"`
	LevelDesc string `yaml:"-"`
	Perm      string `yaml:"perm"`
}

// 蓝星设备接入服务DAServer配置
type DaServer struct {
	Host    string
	Port    int32
	Timeout int32
	Enable  int `yaml:"enable"`
}

func Parse(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("打开配置文件失败：%s", err.Error())
	}

	cfg := Config{}
	y := yaml.NewDecoder(bytes.NewBuffer(data))
	err = y.Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("解析配置文件失败：%s", err.Error())
	}

	Cfg = &cfg
	if cfg.Redis.Enable == 1 {
		if cfg.Redis.Host != "" {
			if Cfg.Redis.HaType == RedisModeOfSignal {
				Cfg.Redis.Hosts = append(Cfg.Redis.Hosts,
					HostInfo{
						Host: Cfg.Redis.Host,
						Port: Cfg.Redis.Port,
					})
			} else {
				Cfg.Redis.Hosts = append(Cfg.Redis.Hosts,
					HostInfo{
						Host: Cfg.Redis.Host,
						Port: Cfg.Redis.Port,
					},
					HostInfo{
						Host: Cfg.Redis.Host1,
						Port: Cfg.Redis.Port1,
					},
					HostInfo{
						Host: Cfg.Redis.Host2,
						Port: Cfg.Redis.Port2,
					})
			}
		}
	}
	return &cfg, nil
}
