package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("notebook", 2000.00)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "notebook", p.Name)
	assert.Equal(t, 2000., p.Price)
}

func TestProductValidateName(t *testing.T) {
	p, err := NewProduct("", 2000.)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductValidatePrice(t *testing.T) {
	p, err := NewProduct("Notebook", .0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductValidateInvalidPrice(t *testing.T) {
	p, err := NewProduct("Notebook", -1)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Notebook", 2000.)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
