package service

import (
	"account-management/root/model"
	"account-management/root/domain/repository"
	"account-management/root/domain/entity"
)


type UserService struct {}

func (us *UserService) SignUpUser(username string, rawPassword string) string {
	// 入力からエンティティを作る
	// エンティティを入力にリポジトリで永続化可能かチェックする(usernameがかぶらないか)
	// エンティティを永続化する
}