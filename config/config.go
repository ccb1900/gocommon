package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var (
	v         *viper.Viper
	logconfig LogConfig
	once      sync.Once
)

type LogRoate struct {
	MaxSize   int  `mapstructure:"size" json:"size"`
	MaxAge    int  `mapstructure:"age" json:"age"`
	MaxBackup int  `mapstructure:"backup" json:"backup"`
	Compress  bool `mapstructure:"compress" json:"compress"`
	LocalTime bool `mapstructure:"localtime" json:"localtime"`
}
type LogConfig struct {
	Path  string    `mapstructure:"path" json:"path"`
	Level string    `mapstructure:"level" json:"level"`
	Roate *LogRoate `mapstructure:"rotate" json:"rotate"`
}

func GetLog() LogConfig {
	return logconfig
}

func Default() *viper.Viper {
	if v == nil {
		panic("config not initialized")
	}
	return v
}

func Init(appname string) {
	once.Do(func() {
		v = viper.New()
	})
	v.SetConfigName("app")               // name of config file (without extension)
	v.SetConfigType("yaml")              // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath("./res")             // path to look for the config file in
	v.AddConfigPath("/etc/" + appname)   // path to look for the config file in
	v.AddConfigPath("$HOME/." + appname) // call multiple times to add many search paths
	v.AddConfigPath(".")                 // optionally look for config in the working directory
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	v.UnmarshalKey("log", &logconfig)
	setDefaultLogConfig(&logconfig)
}

func setDefaultLogConfig(l *LogConfig) {
	if l.Level == "" {
		l.Level = "info"
	}
	if l.Path == "" {
		l.Path = "logs/app.log"
	}

	if l.Roate == nil {
		l.Roate = &LogRoate{
			Compress:  true,
			LocalTime: true,
			MaxSize:   100,
			MaxAge:    30,
			MaxBackup: 10,
		}
	}
}
