package repository

import (
	"context"
	"account-management/root/database"
	"account-management/root/model"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user model.User) error
	FindByUsername(ctx context.Context, username string) (bool, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
}

type UserRepository struct {
	postgresqlInterface database.PostgresqlInterface
}

func NewUserRepository(postgresqlInterface database.PostgresqlInterface) UserRepositoryInterface {
	return &UserRepository{postgresqlInterface: postgresqlInterface}
}

func (userRepository *UserRepository) Create(ctx context.Context, user model.User) error {
	dbConn := userRepository.postgresqlInterface.NewClientConnection()
	defer dbConn.Close()

	d := dbConn.Create(&user)
	if d.Error != nil {
		return d.Error
	}
	dbConn.Save(&user)

	return nil
}

func (userRepository *UserRepository) FindByUsername(ctx context.Context, username string) (bool, error) {
	dbConn := userRepository.postgresqlInterface.NewClientConnection()
	defer dbConn.Close()

	var recordCount = 0

	user := model.User{}
	err := dbConn.Where("username=?", username).First(&user).Count(&recordCount).Error

	if recordCount == 0 {
		return false, err
	} else {
		return true, err
	}
}

func (userRepository *UserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	dbConn := userRepository.postgresqlInterface.NewClientConnection()
	defer dbConn.Close()
	user := model.User{}
	err := dbConn.Where("username=?", username).First(&user).Error

	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}
