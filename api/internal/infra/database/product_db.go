package database

import (
	"github.com/henrymoreirasilva/go-api/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	if sort == "" {
		sort = "asc"
	}

	err := p.DB.Offset((page - 1) * limit).Limit(limit).Order("created_at " + sort).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	product := &entity.Product{}
	err := p.DB.First(product, "id = ?", id).Error
	return product, err
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}

	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}
