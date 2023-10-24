package repository

import (
	"context"
	"time"

	"github.com/vildan-valeev/perx_test/internal/domain"
)

type ItemRepository struct{}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{}
}

func (r *ItemRepository) UpdateCurrentIteration(id int64, currentIteration int) error {
	return nil
}

// SetItem Добавляем item в хранилище!
func (r *ItemRepository) SetItem(ctx context.Context, item *domain.Item) error {
	s := GetLocalStorage()
	return s.SetItem(item)
}

// GetItem Получаем item из хранилища!
func (r *ItemRepository) GetItem(ctx context.Context, id int64) (*domain.Item, error) {
	s := GetLocalStorage()
	return s.GetItem(id)
}

// GetItems Получаем мапу items из хранилища.
func (r *ItemRepository) GetItems(ctx context.Context) (*domain.Items, error) {
	s := GetLocalStorage()
	return s.GetItems()
}

// SetStatus Обновляем поле статус.
func (r *ItemRepository) SetStatus(id int64, status domain.Status) error {
	s := GetLocalStorage()
	return s.SetStatus(id, status)
}

// SetStartTime Обновляем поле время начала обработки таски.
func (r *ItemRepository) SetStartTime(id int64, t time.Time) error {
	s := GetLocalStorage()
	return s.SetStartTime(id, t)
}

// SetEndTime Обновляем поле время окончания обработки таски.
func (r *ItemRepository) SetEndTime(id int64, endTime time.Time) error {
	s := GetLocalStorage()
	return s.SetEndTime(id, endTime)
}

// DeleteItem Удаление записи из мапы.
func (r *ItemRepository) DeleteItem(id int64) error {
	s := GetLocalStorage()
	return s.DeleteItem(id)
}
