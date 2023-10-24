package repository

import (
	"ecom_clean_architecture/domain/model"
	"ecom_clean_architecture/usecase/repo"
	"github.com/jinzhu/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) repo.ProductRepo {
	return &productRepo{db: db}
}

func (p *productRepo) GetAllProducts() ([]*model.Product, error) {
	var products []*model.Product
	err := p.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
