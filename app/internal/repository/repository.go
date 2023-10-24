package repository

import (
	"context"
	"time"

	"github.com/vildan-valeev/perx_test/internal/domain"
)

/*
Работа с БД(памятью)
*/

type Repositories struct {
	Item Item
}

// NewRepositories создаем структуру репозиториев.
func NewRepositories() *Repositories {
	return &Repositories{
		Item: NewItemRepository(),
	}
}

// Item - методы для работы с БД.
type Item interface {
	SetItem(ctx context.Context, item *domain.Item) error
	GetResultCan() <-chan domain.Item
	UpdateCurrentIteration(id int64, currentIteration int) error
	GetItem(ctx context.Context, id int64) (*domain.Item, error)
	GetItems(ctx context.Context) (*domain.Items, error)
	// Update Item
	SetStatus(id int64, status domain.Status) error
	SetStartTime(id int64, time time.Time) error
	SetEndTime(id int64, time time.Time) error
	DeleteItem(id int64) error
}
