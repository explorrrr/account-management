package entity_test

import (
	"math/rand"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
	"account-management/root/domain/entity"
)

// パスワードのバリデーションテスト
func TestValidateInputPassword(t *testing.T) {
	const lowerThreshold = 8
	const higherThreshold = 128

	// 7文字の場合
	resultLessCharacters, _ := entity.ValidateInputPassword(strings.Repeat("a", lowerThreshold - 1))
	// 129文字の場合
	resultGreaterCharacters, _ := entity.ValidateInputPassword(strings.Repeat("a", higherThreshold + 1))
	// 8 ~ 128文字の場合(どれでも成功するはずなので8~128文字でランダムにテスト)
	successResult, _ := entity.ValidateInputPassword(strings.Repeat("a", lowerThreshold + rand.Intn(higherThreshold-lowerThreshold+1)))
	// 数値の場合のテスト(文字列)
	successNumberResult, _ := entity.ValidateInputPassword(strings.Repeat("1", lowerThreshold + rand.Intn(higherThreshold-lowerThreshold+1)))
	// 記号の場合のテスト(文字列)
	successSymbolResult, _ := entity.ValidateInputPassword(strings.Repeat("+", lowerThreshold + rand.Intn(higherThreshold-lowerThreshold+1)))

	assert.Equal(t, resultGreaterCharacters, false)
	assert.Equal(t, resultLessCharacters, false)
	assert.Equal(t, successResult, true)
	assert.Equal(t, successNumberResult, true)
	assert.Equal(t, successSymbolResult, true)

}

// ユーザー名のバリデーションテスト
func TestValidateInputUsername(t *testing.T) {
	const lowerThreshold = 6
	const higherThreshold = 64

	// 5文字の場合
	resultLessCharacters, _ := entity.ValidateInputUsername(strings.Repeat("a", lowerThreshold - 1))
	// 65文字の場合
	resultGreaterCharacters, _ := entity.ValidateInputUsername(strings.Repeat("a", higherThreshold + 1))

	// 1 ~ 64文字までの場合
	successResult, _ := entity.ValidateInputUsername(strings.Repeat("a", lowerThreshold + rand.Intn(higherThreshold-lowerThreshold+1)))

	assert.Equal(t, resultLessCharacters, false)
	assert.Equal(t, resultGreaterCharacters, false)
	assert.Equal(t, successResult, true)

}

// ユーザーが正常のnewできるか
func TestUserCreate(t *testing.T) {
	_, err := entity.NewUser("test_user", "test_password")
	if err != nil {
		t.Fatal(err.Error())
	}
}
