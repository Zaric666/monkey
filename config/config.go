package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Application Application
	Databases   map[string]*Database
	Log         Log
}

type Application struct {
	Mode          string
	Name          string
	Host          string
	Port          int
	ReadTimeout   int
	WriterTimeout int
	EnableDp      bool
}

// MySQL相关配置
type Database struct {
	Driver string `mapstructure:"driver"`
	Source string `mapstructure:"source"`
}

// 日志保存地址
type Log struct {
	Path  string
	Level string
}

var c Config

func init() {
	// 设置文件名
	viper.SetConfigName("settings")
	// 设置文件类型
	viper.SetConfigType("yaml")
	// 设置文件路径，可以多个viper会根据设置顺序依次查找
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	viper.Unmarshal(&c)
}

func GetConfig() Config {
	return c
}
