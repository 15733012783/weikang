package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func InitViper(fileName, filePath, configName string) (map[string]string, error) {
	viper.SetConfigName(fileName)
	viper.AddConfigPath(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	mapString := viper.GetStringMapString(configName)
	return mapString, nil
}

type BlogBff struct {
	//ServiceName string       `yaml:"serverName"`
	Nacos NacosConfig `mapstructure:"nacosConfig"`
}
type NacosConfig struct {
	ServerConfig struct {
		IpAddr string `mapstructure:"IpAddr"`
		Port   uint64 `mapstructure:"Port"`
	}
	ClientConfig struct {
		NamespaceId string `mapstructure:"NamespaceId"`
	}
	ConfigParam struct {
		DataId string `mapstructure:"DataId"`
		Group  string `mapstructure:"Group"`
	}
}

var NacosConf NacosConfig

func InitViperByStruct(configFile string) error {
	v := viper.New()
	v.SetConfigFile(configFile)

	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	err = v.Unmarshal(&NacosConf)
	if err != nil {
		return err
	}

	v.WatchConfig()
	//动态监听

	v.OnConfigChange(func(c fsnotify.Event) {
		err = v.Unmarshal(&NacosConf)
		if err != nil {
			return
		}
		log.Println("本地配置文件发生变动")
	})
	log.Println("viper初始化完成")
	return nil
}
