package service

import (
	"context"

	"github.com/vildan-valeev/perx_test/internal/domain"
	"github.com/vildan-valeev/perx_test/internal/repository"
	"github.com/vildan-valeev/perx_test/internal/transport/dto"
)

/*
Работа в бизнес логикой
*/

type Item interface {
	AddItemToQueueService(ctx context.Context, item *dto.ItemToQueueDTO) error
	ListItemService(ctx context.Context) (*domain.Items, error)
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
