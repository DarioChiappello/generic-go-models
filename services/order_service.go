package services

import (
	"github.com/DarioChiappello/generic-go-models/models"
	"github.com/DarioChiappello/generic-go-models/repositories"

	"gorm.io/gorm"
)

type OrderService struct {
	*BaseService[models.Order]
}

func NewOrderService(db *gorm.DB) *OrderService {
	repo := repositories.NewBaseRepository[models.Order](db)
	return &OrderService{
		BaseService: NewBaseService(repo),
	}
}
