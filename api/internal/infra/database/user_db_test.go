package database

import (
	"testing"

	"github.com/henrymoreirasilva/go-api/internal/entity"
	"github.com/stretchr/testify/assert"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(entity.User{})
	user, _ := entity.NewUser("henry", "henry@zoomwi.com.br", "1234")

	uDB := NewUser(db)
	err = uDB.Create(user)
	if err != nil {
		t.Error(err)
	}

	var userFound entity.User

	err = uDB.DB.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, userFound.Name, user.Name)
	assert.Equal(t, userFound.Email, user.Email)
	assert.NotNil(t, userFound.Password)
}

func TestUserFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"))
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(entity.User{})
	user, err := entity.NewUser("henry", "henry@zoomwi.com.br", "12345")
	if err != nil {
		t.Error(err)
	}

	userDB := NewUser(db)
	err = userDB.Create(user)
	assert.Nil(t, err)

	userEmail, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, userEmail.Email, user.Email)

}
