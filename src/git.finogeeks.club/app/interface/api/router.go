package api

import (
	"git.finogeeks.club/app/interface/api/handler"

	"github.com/labstack/echo"
)

func Load(e *echo.Echo) {
	userHandler := handler.NewUserHandler()
	baseGroup := e.Group("/api/v1/echo")
	{
		baseGroup.GET("/user", userHandler.ListUser)
		baseGroup.POST("/register", userHandler.Register)

	}
}
