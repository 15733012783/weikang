package config

import (
	"github.com/spf13/viper"
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
