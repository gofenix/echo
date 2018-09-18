package usecase

import (
	"fmt"
	"log"

	"git.finogeeks.club/app/domain/model"
	"git.finogeeks.club/app/domain/repository"
	"git.finogeeks.club/app/domain/service"
	"github.com/google/uuid"
)

type UserUsecase interface {
	ListUser() ([]*User, error)
	RegisterUser(email string) error
	GetTest()
}

type userUsecase struct {
	repo    repository.UserRepository
	service *service.UserService
}

func NewUserUsecase(repo repository.UserRepository, service *service.UserService) UserUsecase {
	return &userUsecase{
		repo:    repo,
		service: service,
	}
}

func (u *userUsecase) GetTest() {
	fmt.Println("uuuuuu")

}

func (u *userUsecase) ListUser() ([]*User, error) {
	fmt.Println("aaaaaaaaaaaaaaaa")
	log.Println(&u)
	log.Println(&u.repo)
	users, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return toUser(users), nil
}

func (u *userUsecase) RegisterUser(email string) error {
	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	if err := u.service.Duplicated(email); err != nil {
		return err
	}
	user := model.NewUser(uid.String(), email)
	if err := u.repo.Save(user); err != nil {
		return err
	}
	return nil
}

type User struct {
	ID    string
	Email string
}

func toUser(users []*model.User) []*User {
	res := make([]*User, len(users))
	for i, user := range users {
		res[i] = &User{
			ID:    user.GetID(),
			Email: user.GetEmail(),
		}
	}
	return res
}
