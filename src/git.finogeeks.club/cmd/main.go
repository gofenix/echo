package main

import (
	"fmt"
	"os"

	"git.finogeeks.club/app/interface/api"
	"git.finogeeks.club/app/interface/persistence/mongodb"

	"github.com/labstack/echo"

	"git.finogeeks.club/config"

	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()
	mongodb.Init()

	fmt.Println("hello main")
	configExample()

	e := echo.New()
	api.Load(e)

	//fmt.Println(http.ListenAndServe(":"+viper.GetString("http.port"), e).Error())
	e.Start(":" + viper.GetString("http.port"))
}

func configExample() {
	fmt.Println(viper.GetString("url"))
	os.Setenv("URL", "888888")
	fmt.Println(os.Getenv("URL"))
	fmt.Println(viper.GetString("url"))

	fmt.Println(viper.GetString("http.port"))
	os.Setenv("HTTP_PORT", "8888")
	fmt.Println(os.Getenv("HTTP_PORT"))
	fmt.Println(viper.GetString("http.port"))
}
