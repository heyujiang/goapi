package config

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	fmt.Println(cfg)

	config := Config{
		Name: cfg,
	}

	if err := config.initConfig(); err != nil {
		return errors.New("Config initialize error : " + err.Error())
	}

	config.watchConfig()

	return nil
}

func (config *Config) initConfig() error {

	if config.Name != "" {
		viper.SetConfigFile(config.Name) //如果指定了配置文件，解析指定的文件
	} else {
		viper.AddConfigPath("/Users/fangyamin/develop/golang/goapi/conf/") //解析默认的
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("GOAPI")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (config *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Config is changed : %s", in.Name)
	})
}
