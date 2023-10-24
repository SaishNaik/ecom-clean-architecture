package dependency_injection

import (
	"ecom_clean_architecture/adapter/controller"
	"ecom_clean_architecture/adapter/repository"
	"ecom_clean_architecture/usecase/usecase"
	"github.com/jinzhu/gorm"
)

type depInjection struct {
	db *gorm.DB
}

type DepInjection interface {
	NewAppController() controller.AppController
}

func NewDepInjection(db *gorm.DB) DepInjection {
	return &depInjection{db}
}

func (d *depInjection) NewAppController() controller.AppController {
	productRepo := repository.NewProductRepo(d.db)
	product := usecase.NewProduct(productRepo)

	return controller.AppController{
		Product: controller.NewProductController(product),
	}
}
