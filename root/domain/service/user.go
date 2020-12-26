package service

import (
	"context"
	"errors"
	"account-management/root/model"
	"account-management/root/database"
	"account-management/root/domain/repository"
	"account-management/root/domain/entity"
)


type UserService struct {
	UserRepositoryInterface repository.UserRepositoryInterface
}

func (us *UserService) SignUpUser(ctx context.Context, username string, rawPassword string) (bool, error) {
	// FIXME どこでDI設定するべきか考える
	dataStoreInterface := database.NewPostgresql()
	us.UserRepositoryInterface = repository.NewUserRepository(dataStoreInterface)

	// 入力からエンティティを作る
	userEntity, err := entity.NewUser(username, rawPassword)
	if err != nil {
		return false, err
	}

	userModel := model.User{Username: userEntity.Username, Password: userEntity.Password}
	// エンティティを入力にリポジトリで永続化可能かチェックする(usernameがかぶらないか)
	exists, _ := us.UserRepositoryInterface.FindByUsername(ctx, userEntity.Username)

	if exists == true {
		return false, errors.New("User already exist, please try another username")
	} else {
		// エンティティを永続化する
		us.UserRepositoryInterface.Create(ctx, userModel)
		return true, nil
	}
}
