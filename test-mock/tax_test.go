package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", .0).Return(errors.New("erro ao salvar taxa"))

	err := CalculateTaxAndSave(1000.00, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(.00, repository)
	assert.Error(t, err)

	repository.AssertExpectations(t)
}
