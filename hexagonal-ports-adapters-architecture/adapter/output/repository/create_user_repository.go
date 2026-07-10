package repository

import (
	"context"
	"golang_hexagonal_architecture/adapter/output/converter"
	"golang_hexagonal_architecture/adapter/output/model"
	"golang_hexagonal_architecture/adapter/output/table"
	"golang_hexagonal_architecture/application/domain"
	"golang_hexagonal_architecture/application/port/output"
	"golang_hexagonal_architecture/configuration/logger"
	"time"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func NewUserRepository(database *pgx.Conn) output.UserPort {
	return &userRepository{database}
}

type userRepository struct {
	databaseConnection *pgx.Conn
}

func (ur *userRepository) CreateUser(userDomain domain.User) (*domain.User, error) {
	logger.Info("Init createUser repository", zap.String("journey", "createUser"))

	now := time.Now()

	value := converter.ConvertDomainToEntity(userDomain)

	stmt := table.Users.INSERT(
		table.Users.FirstName,
		table.Users.LastName,
		table.Users.Age,
		table.Users.CreatedAt,
		table.Users.UpdatedAt,
	).VALUES(
		value.FirstName,
		value.LastName,
		value.Age,
		now,
		now,
	).RETURNING(table.Users.AllColumns)

	query, args := stmt.Sql()

	var result model.Users
	err := ur.databaseConnection.QueryRow(context.Background(), query, args...).Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.Age,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)
	if err != nil {
		return nil, nil
	}

	value.ID = result.ID
	
	return converter.ConvertEntityToDomain(*value), nil
}
