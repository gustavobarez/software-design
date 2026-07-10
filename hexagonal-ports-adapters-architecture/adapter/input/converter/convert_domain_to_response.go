package converter

import (
	"golang_hexagonal_architecture/adapter/input/model/response"
	"golang_hexagonal_architecture/application/domain"
)

func ConvertDomainToResponse(userDomain *domain.User) response.UserResponse {
	return response.UserResponse{
		ID: userDomain.ID,
		FirstName: userDomain.FirstName,
		LastName: userDomain.LastName,
		Age: userDomain.Age,
	}
}