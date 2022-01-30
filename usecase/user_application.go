package usecase

import (
	"github.com/google/uuid"
	"log"
	"prac-orm-transaction/domain/repository"
)

type UserUseCase interface {
	CreateUserAccount(id, pass string) (string, error)
	SignIn(id, pass string) (string, error)
}

type userUseCase struct {
	user repository.UserRepository
}

func NewUserUseCase(user repository.UserRepository) *userUseCase {
	return &userUseCase{user: user}
}

func (uu userUseCase) CreateUserAccount(id, pass string) (string, error) {
	var token string

	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return token, err
	}

	token = uuid.String()

	if err = uu.user.CreateUsersAccount(id, pass, token); err != nil {
		log.Println(err)
		return token, err
	}

	return token, nil
}

func (uu userUseCase) SignIn(id, pass string) (string, error) {
	var token string

	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return token, err
	}

	token = uuid.String()

	if err = uu.user.RegisterUsersInfo(id, pass, token); err != nil {
		log.Println(err)
		return token, err
	}

	return token, nil
}
