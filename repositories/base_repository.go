package repositories

import (
	"reflect"

	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) GetAll() ([]T, error) {
	var models []T
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (r *BaseRepository[T]) GetByID(id uint) (*T, error) {
	var model T
	if err := r.db.First(&model, id).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *BaseRepository[T]) Create(model *T) error {
	return r.db.Create(model).Error
}

func (r *BaseRepository[T]) Update(model *T) error {
	return r.db.Model(model).
		Where("id = ?", getModelID(model)).
		Updates(model).
		Error
}

func getModelID[T any](model *T) uint {
	val := reflect.ValueOf(model).Elem()
	idField := val.FieldByName("ID")
	if !idField.IsValid() {
		panic("Model doesn't have field ID")
	}
	return uint(idField.Uint())
}

func (r *BaseRepository[T]) Patch(id uint, updates map[string]interface{}) error {
	return r.db.Model(new(T)).
		Where("id = ?", id).
		Updates(updates).
		Error
}

func (r *BaseRepository[T]) Delete(id uint) error {
	return r.db.Model(new(T)).
		Where("id = ?", id).
		Delete(new(T)).
		Error
}
