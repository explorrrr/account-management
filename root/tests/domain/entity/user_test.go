package entity_test

import (
	"testing"
	"account-management/root/domain/entity"
)

func TestUserCreate(t *testing.T) {
	_, err := entity.NewUser("test", "test_password")
	if err != nil {
		t.Fatal(err.Error())
	}
}
