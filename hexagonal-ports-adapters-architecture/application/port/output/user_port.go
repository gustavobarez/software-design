package output

import "golang_hexagonal_architecture/application/domain"

type UserPort interface {
	CreateUser(domain.User) (*domain.User, error)
}