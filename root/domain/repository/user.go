package repository

import (
	"context"
	"account-management/root/model"
	"account-management/root/database"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *model.User) error
	// Find(ctx context.Context, id int) (bool, error)
}

type UserRepository struct {
	postgresqlinterface database.PostgresqlInterface
}

func NewUserRepository(postgresqlinterface database.PostgresqlInterface) UserRepositoryInterface {
	return &UserRepository{postgresqlinterface: postgresqlinterface}
}

func (UserRepository *UserRepository) Create(ctx context.Context, user model.User) (*model.User, error) {
	dbConn := UserRepository.postgresqlinterface.NewClientConnection()
	defer dbConn.Close()

	d := dbConn.Create(&user)
	if d.Error != nil {
		return nil, d.Error
	}

	return &user, nil
}
