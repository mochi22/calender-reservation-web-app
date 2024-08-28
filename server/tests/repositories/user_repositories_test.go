package repositories_test

import (
	"testing"	
	"github.com/mochi22/calender-reservation-web-app/server/internals/models"
	"github.com/mochi22/calender-reservation-web-app/server/internals/repositories"

	"github.com/stretchr/testify/assert"
)


func TestUserRepository_Create(t *testing.T) {
	// テストデータベースに接続
	db, teardown := testDB()
	defer teardown()

	repo := repositories.NewUserRepository(db)

	user := &models.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
		Role:     "user",
	}

	err := repo.Create(user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	// テストデータベースに接続
	db, teardown := testDB()
	defer teardown()

	repo := repositories.NewUserRepository(db)

	// テストユーザーを作成
	user := &models.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
		Role:     "user",
	}

	err := repo.Create(user)
	assert.NoError(t, err)

	// メールアドレスでユーザーを検索
	foundUser, err := repo.FindByEmail("test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, user.ID, foundUser.ID)
	assert.Equal(t, user.Name, foundUser.Name)
	assert.Equal(t, user.Email, foundUser.Email)
	assert.Equal(t, user.Password, foundUser.Password)
	assert.Equal(t, user.Role, foundUser.Role)
}