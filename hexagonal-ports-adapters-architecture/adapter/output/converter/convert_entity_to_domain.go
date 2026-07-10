package converter

import (
	"golang_hexagonal_architecture/adapter/output/model"
	"golang_hexagonal_architecture/application/domain"
)

func ConvertEntityToDomain(user model.Users) (*domain.User) {
	domainConverted := &domain.User{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Age: user.Age,
	}

	domainConverted.ID = user.ID
	return domainConverted
}