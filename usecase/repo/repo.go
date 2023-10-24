package repo

import "ecom_clean_architecture/domain/model"

type ProductRepo interface {
	GetAllProducts() ([]*model.Product, error)
}
