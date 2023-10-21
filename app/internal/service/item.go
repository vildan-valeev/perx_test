package service

import (
	"context"
	"crypto/rand"
	"log"
	"math/big"

	"github.com/vildan-valeev/perx_test/internal/domain"
	"github.com/vildan-valeev/perx_test/internal/repository"
	"github.com/vildan-valeev/perx_test/internal/transport/dto"
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

	n, err := rand.Int(rand.Reader, big.NewInt(27))
	if err != nil {
		log.Fatal(err)
	}

	i := domain.Item{
		ID:            n.Int64(),
		ElementsCount: addItem.N,
		Delta:         addItem.D,
		StartElement:  addItem.N1,
		TimeInterval:  addItem.I,
		TTL:           addItem.TTL,
	}

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

	return resp, nil
}
