package handler

import (
	"fmt"
	"net/http"

	"git.finogeeks.club/app/domain/service"

	"git.finogeeks.club/app/interface/persistence/mongodb"

	"git.finogeeks.club/app/usecase"
	"github.com/labstack/echo"
)

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler() *userHandler {
	var repo = mongodb.NewUserRepository()
	fmt.Println("repo----->", repo)
	svc := service.NewUserService(repo)
	var userUsecase = usecase.NewUserUsecase(repo, svc)
	return &userHandler{
		userUsecase: userUsecase,
	}
}

func (h *userHandler) GetTest(c echo.Context) error {
	fmt.Println("-----")
	fmt.Println(h.userUsecase)
	h.userUsecase.GetTest()

	return c.String(200, "yes you are right")
}

func (h *userHandler) ListUser(c echo.Context) error {
	fmt.Println("-----")
	fmt.Println(h.userUsecase)
	users, err := h.userUsecase.ListUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) Register(c echo.Context) error {
	email := c.Param("email")
	err := h.userUsecase.RegisterUser(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]string{})
}
