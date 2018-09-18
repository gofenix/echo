package main

import (
	"fmt"
	"os"

	"git.finogeeks.club/config"
	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()

	fmt.Println("hello main")
	configExample()
}

func configExample() {
	fmt.Println(viper.GetString("url"))
	os.Setenv("URL", "888888")
	fmt.Println(os.Getenv("URL"))
	fmt.Println(viper.GetString("url"))

	fmt.Println(viper.GetString("http.port"))
	os.Setenv("HTTP_PORT", "888888")
	fmt.Println(os.Getenv("HTTP_PORT"))
	fmt.Println(viper.GetString("http.port"))
}
