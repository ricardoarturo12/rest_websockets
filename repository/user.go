package repository

import (
	"context"

	"github.com/ricardoarturo12/rest_websockets/models"
)

// codigo basado en abstracciones y no en cosas concretas

//abstracciones
//Handler - GetUserByIdPostgres ......
//Handler - GetUserByIdMongo ....const
// Handler - GetUserById - User
// Postgres
// GetUserByIdMongo..

//concretas

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	Close() error
}

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func Close() error {
	return implementation.Close()
}
