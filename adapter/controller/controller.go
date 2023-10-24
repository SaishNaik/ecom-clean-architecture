package controller

import (
	"ecom_clean_architecture/usecase/usecase"
	"net/http"
)

type ProductController interface {
	GetAllProducts(ctx Context) error
}

type productController struct {
	Product usecase.Product
}

func NewProductController(product usecase.Product) ProductController {
	return &productController{Product: product}
}

func (c *productController) GetAllProducts(ctx Context) error {
	// get params
	// no need of params

	// call usecase
	products, err := c.Product.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &products)
		return err
	}

	ctx.JSON(http.StatusOK, &products)
	return nil
}
