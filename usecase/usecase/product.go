package usecase

import (
	"ecom_clean_architecture/domain/model"
	"ecom_clean_architecture/usecase/repo"
)

type Product interface {
	GetAllProducts() ([]*model.Product, error)
}

type product struct {
	Repo repo.ProductRepo
}

func NewProduct(repo repo.ProductRepo) Product {
	return &product{Repo: repo}
}

func (p *product) GetAllProducts() ([]*model.Product, error) {
	return p.Repo.GetAllProducts()
}
