package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserr(t *testing.T) {
	user, err := NewUser("henry", "henry@zoomwi.com.br", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "henry", user.Name)
	assert.Equal(t, "henry@zoomwi.com.br", user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("henry", "henry@zoomwi.com.br", "123456")
	assert.Nil(t, err)
	assert.NotEqual(t, "123456", user.Password)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))

}
