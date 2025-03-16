package services

import (
	"github.com/DarioChiappello/generic-go-models/models"
	"github.com/DarioChiappello/generic-go-models/repositories"

	"gorm.io/gorm"
)

type UserService struct {
	*BaseService[models.User]
}

func NewUserService(db *gorm.DB) *UserService {
	repo := repositories.NewBaseRepository[models.User](db)
	return &UserService{
		BaseService: NewBaseService(repo),
	}
}
