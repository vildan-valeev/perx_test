package service

import (
	"context"
	"log"
	"math/rand"

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

// AddItemToQueueService Добавление задачи в очеред.
func (c ItemService) AddItemToQueueService(ctx context.Context, addItem *dto.ItemToQueueDTO) error {
	log.Println("Add to queue", addItem)
	i := domain.Item{
		ID:            rand.Int(),
		ElementsCount: addItem.N,
		Delta:         addItem.D,
		StartElement:  addItem.N1,
		TimeInterval:  addItem.I,
		TTL:           addItem.TTL,
	}
	t := i.EndTime.Unix()
	// todo: validation Item

	if err := c.repo.AddTaskRepo(ctx, &i); err != nil {
		return err
	}

	return nil
}

// ListItemService Получение списка.
func (c ItemService) ListItemService(ctx context.Context) (domain.Items, error) {
	resp := domain.Items{
		{
			ID: 0,
		},
		{
			ID: 1,
		},
		{
			ID: 2,
		},
	}
	//c.repo.ListItemRepo(ctx)

	return resp, nil
}
