package repository

import (
	"context"

	"perx/internal/domain"
)

type ItemRepository struct{}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{}
}

func (s *ItemRepository) AddTaskRepo(ctx context.Context, item *domain.Item) error {
	return nil
}

func (s *ItemRepository) ListItemRepo(ctx context.Context) ([]*domain.Item, error) {
	items := make([]*domain.Item, 0)
	return items, nil
}
