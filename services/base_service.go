package services

import "github.com/DarioChiappello/generic-go-models/repositories"

type BaseService[T any] struct {
	repo *repositories.BaseRepository[T]
}

func NewBaseService[T any](repo *repositories.BaseRepository[T]) *BaseService[T] {
	return &BaseService[T]{repo: repo}
}

func (s *BaseService[T]) GetAll() ([]T, error) {
	return s.repo.GetAll()
}

func (s *BaseService[T]) GetByID(id uint) (*T, error) {
	return s.repo.GetByID(id)
}

func (s *BaseService[T]) Create(model *T) error {
	return s.repo.Create(model)
}

func (s *BaseService[T]) Update(model *T) error {
	return s.repo.Update(model)
}

func (s *BaseService[T]) Patch(id uint, updates map[string]interface{}) error {
	return s.repo.Patch(id, updates)
}

func (s *BaseService[T]) Delete(id uint) error {
	return s.repo.Delete(id)
}
