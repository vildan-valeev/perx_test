package service

import (
	"context"

	"perx/internal/domain"
	"perx/internal/repository"
	"perx/internal/transport/dto"
)

// ItemService UseCase - бизнес логика.
type ItemService struct {
	repo repository.Item
}

func NewItemService(repo repository.Item) *ItemService {
	return &ItemService{
		repo: repo,
	}
}

// CreateItemService Создание.
func (c ItemService) CreateItemService(ctx context.Context, itemDTO *dto.ItemDTO) error {
	item := domain.Item{
		ID: 0,
	}

	if err := c.repo.InsertItemRepo(ctx, &item); err != nil {
		return err
	}

	return nil
}

// ListItemService GetAll Получение списка.
func (c ItemService) ListItemService(ctx context.Context) ([]*domain.Item, error) {
	return c.repo.ListItemRepo(ctx)
}
