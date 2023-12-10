package config

import (
	"github.com/spf13/viper"
)

func Loader(path string) {
	// viper.AddConfigPath(".")
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func GetString(param string) string {
	// fmt.Println(param, viper.GetString(param))
	return viper.GetString(param)
}
