package service

import (
	"golang_hexagonal_architecture/application/domain"
	"golang_hexagonal_architecture/application/port/output"
	"golang_hexagonal_architecture/configuration/logger"
)

type userService struct {
	userPort output.UserPort
}

func NewUserService(userPort output.UserPort) *userService {
	return &userService{userPort}
}

func (us *userService) CreateUserService(userDomain domain.User) (*domain.User, error) {
	logger.Info("Init createUserService function")

	userDomainResponse, err := us.userPort.CreateUser(userDomain)
	return userDomainResponse, err

}
