package service

import (
	"context"

	"perx/internal/domain"
	"perx/internal/repository"
	"perx/internal/transport/dto"
)

/*
Работа в бизнес логикой
*/

type Item interface {
	CreateItemService(ctx context.Context, item *dto.ItemDTO) error
	ListItemService(ctx context.Context) ([]*domain.Item, error)
}

type Services struct {
	Item Item
}

type Deps struct {
	Repos *repository.Repositories
	Host  string
}

func NewServices(deps Deps) *Services {
	itemService := NewItemService(deps.Repos.Item)

	return &Services{
		Item: itemService,
	}
}
