package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/henrymoreirasilva/go-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestProductCreate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"))
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("notebook", 2200.00)
	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestProductFindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"))
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("product 1", 10.1)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.Nil(t, err)

	productFind, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.Name, productFind.Name)
	assert.Equal(t, product.Price, productFind.Price)

}

func TestProductFindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"))
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	for x := 1; x < 21; x++ {
		product, _ := entity.NewProduct(fmt.Sprintf("Product %d", x), rand.Float64()*100)
		err = productDB.Create(product)
		assert.Nil(t, err)
	}

	products, err := productDB.FindAll(0, 10, "asc")
	assert.Nil(t, err)

	assert.Len(t, products, 10)
	assert.Equal(t, products[0].Name, "Product 1")
	assert.Equal(t, products[9].Name, "Product 10")

	products, err = productDB.FindAll(2, 10, "asc")
	assert.Nil(t, err)

	assert.Len(t, products, 10)
	assert.Equal(t, products[0].Name, "Product 11")
	assert.Equal(t, products[9].Name, "Product 20")
}

func TestProductDelet(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"))
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("product 1", 10.1)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.Nil(t, err)

	id := product.ID

	err = productDB.Delete(id.String())
	assert.Nil(t, err)

	productFind, err := productDB.FindByID(id.String())
	assert.Error(t, err)
	assert.Equal(t, productFind, &entity.Product{})
}

func TestProductUpdate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"))
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("product 1", 10.1)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.Nil(t, err)

	product.Name = "Produto alterado"
	err = productDB.Update(product)
	assert.Nil(t, err)

	productFind, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, productFind.Name, product.Name)
	assert.Equal(t, productFind.ID, product.ID)
	assert.NotEqual(t, productFind.Name, "product 1")
}
