package input

import "golang_hexagonal_architecture/application/domain"

type UserDomainService interface {
	CreateUserService(
		domain.User,
	) (*domain.User, error)
}