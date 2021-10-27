package config

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type config struct {
	ConfigFile string
}

func Init(configFile string) error {
	fmt.Println(configFile)

	config := config{
		ConfigFile: configFile,
	}

	if err := config.initConfig(); err != nil { //初始化配置文件失败
		return errors.New("Config initialize error : " + err.Error())
	}

	config.watchConfig()

	return nil
}

//初始化配置文件信息
func (config *config) initConfig() error {
	if config.ConfigFile != "" {
		viper.SetConfigFile(config.ConfigFile) //如果指定了配置文件，解析指定的文件
	} else {
		//viper.AddConfigPath("/Users/fangyamin/develop/golang/goapi/config/") //解析默认的
		viper.AddConfigPath("/develop/golang/goapi/config/") //解析默认的
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

//检测配置文件变化
func (config *config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Config is changed : %s", in.Name)
	})
}
