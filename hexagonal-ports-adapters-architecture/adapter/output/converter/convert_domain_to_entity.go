package converter

import (
	"golang_hexagonal_architecture/adapter/output/model"
	"golang_hexagonal_architecture/application/domain"
)

func ConvertDomainToEntity(user domain.User) (*model.Users) {
	return &model.Users{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Age: user.Age,
	}
}