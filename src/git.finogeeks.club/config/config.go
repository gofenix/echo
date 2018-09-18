package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig() {

	viper.AddConfigPath("$GOPATH/src/git.finogeeks.club/config")
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("something error")
		panic(err)
	}
}
