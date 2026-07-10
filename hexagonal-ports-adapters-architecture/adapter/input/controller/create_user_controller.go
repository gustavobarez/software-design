package controller

import (
	"encoding/json"
	"golang_hexagonal_architecture/adapter/input/converter"
	"golang_hexagonal_architecture/adapter/input/model/request"
	"golang_hexagonal_architecture/application/domain"
	"golang_hexagonal_architecture/application/port/input"
	"golang_hexagonal_architecture/configuration/logger"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

func NewUserControllerInterface(serviceInterface input.UserDomainService) UserControllerInterface {
	return &userControllerInterface{service: serviceInterface}
}

type UserControllerInterface interface {
	CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type userControllerInterface struct {
	service input.UserDomainService
}

func (uc *userControllerInterface) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logger.Info("Init createUser controller", zap.String("journey", "createUser"))

	var userRequest request.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		return
	}

	userDomain := domain.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Age:       userRequest.Age,
	}
	domainResult, err := uc.service.CreateUserService(userDomain)
	if err != nil {
		return
	}

	logger.Info("createUser executed successfully", zap.Int32("userId", domainResult.ID), zap.String("journey", "createUser"))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(converter.ConvertDomainToResponse(domainResult)); err != nil {
		return
	}
}
