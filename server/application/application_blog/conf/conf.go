package conf

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	ContextReqUUid    = "uuid"
	ContextSourceName = "source"
	GrpcTimeOut       = 5 * time.Second
)

var Cfg *config

type config struct {
	Log     Log        `yaml:"logger"`
	Service ServiceCfg `yaml:"service"`
	Server  Server     `yaml:"server"`
}

type Server struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}

type HostInfo struct {
	Host string
	Port int
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

type ServiceCfg struct {
	GrpcIp   string `yaml:"grpcIp"`
	GrpcPort int    `yaml:"grpcPort"`
}

func Parse(filename string) (*config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("打开配置文件失败：%s", err.Error())
	}

	cfg := config{}
	y := yaml.NewDecoder(bytes.NewBuffer(data))
	err = y.Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("解析配置文件失败：%s", err.Error())
	}

	Cfg = &cfg

	return &cfg, nil
}
