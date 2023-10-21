package repository

import (
	"context"

	"github.com/vildan-valeev/perx_test/internal/domain"
)

/*
Работа с БД
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
	AddTaskRepo(ctx context.Context, i *domain.Item) error
	ListItemRepo(ctx context.Context) ([]*domain.Item, error)
}
