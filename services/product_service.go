package services

import (
	"github.com/DarioChiappello/generic-go-models/models"
	"github.com/DarioChiappello/generic-go-models/repositories"

	"gorm.io/gorm"
)

type ProductService struct {
	*BaseService[models.Product]
}

func NewProductService(db *gorm.DB) *ProductService {
	repo := repositories.NewBaseRepository[models.Product](db)
	return &ProductService{
		BaseService: NewBaseService(repo),
	}
}
