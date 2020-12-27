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

func(us *UserService) AuthUser(ctx context.Context, username string, rawPassword string) (*entity.JWT, error) {
	// FIXME どこでDI設定するべきか考える
	dataStoreInterface := database.NewPostgresql()
	us.UserRepositoryInterface = repository.NewUserRepository(dataStoreInterface)

	// Userオブジェクトの取得
	user, err := us.UserRepositoryInterface.GetByUsername(ctx, username)

	if err != nil {
		return nil, err
	}

	userEntity := &entity.UserEntity{Username: user.Username, Password: user.Password}

	token, err := userEntity.IssueJWT(rawPassword)

	if err != nil {
		return nil, err
	}

	jwt := &entity.JWT{Token: token}

	return jwt, nil
}

func(us *UserService) ValidateToken(ctx context.Context, token string) (string, error) {
	tokenEntity := &entity.JWT{Token: token}
	code, err := tokenEntity.Validate()

	return code, err
}

func(us *UserService) ChangePassword(ctx context.Context, username string, currentPassword string, desiredPassword string) (bool, string, error) {
	// FIXME どこでDI設定するべきか考える
	dataStoreInterface := database.NewPostgresql()
	us.UserRepositoryInterface = repository.NewUserRepository(dataStoreInterface)

	// Userオブジェクトの取得
	user, err := us.UserRepositoryInterface.GetByUsername(ctx, username)

	if err != nil {
		return false, "9999", err
	}

	userEntity := &entity.UserEntity{Username: user.Username, Password: user.Password}

	userEntity, err = userEntity.ChangePassword(currentPassword, desiredPassword)

	if err != nil {
		return false, "9999", err
	}

	_ = us.UserRepositoryInterface.Update(ctx, username, map[string]interface{}{"password": userEntity.Password})

	return true, "0000", nil

}
