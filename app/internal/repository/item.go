package repository

import (
	"context"
	"github.com/vildan-valeev/perx_test/internal/domain"
	"log"
	"sync"
)

type itemRepository struct {
	items    domain.Items // токен от Транзита
	apiKey   string       // Ключ от Payment
	resultCh chan<- domain.Item
	sync.RWMutex
}

var instance *itemRepository //nolint:gochecknoglobals

var once sync.Once //nolint:gochecknoglobals

func NewItemRepository() Item {
	once.Do(func() {
		instance = new(itemRepository)

		log.Println("LocalStorage init...")
	})

	return instance
}

func (r *itemRepository) UpdateCurrentIteration(id int64, currentIteration int) error {
	r.Lock()
	defer r.Unlock()
	log.Println("Обновление текущей итерации  item!")

	r.items[id].CurrentIteration = currentIteration
	return nil
}

func (r *itemRepository) GetResultCan() chan<- domain.Item {
	r.Lock()
	defer r.Unlock()
	log.Println("запрашиваем канал для посылки результатов!")

	return r.resultCh
}

// SetItem Добавляем item в хранилище!
func (r *itemRepository) SetItem(ctx context.Context, item *domain.Item) error {
	r.Lock()
	defer r.Unlock()
	log.Printf("Добавляем item %d в хранилище! \n", item.ID)

	r.items[item.ID] = item
	return nil
}

// GetItem Получаем item из хранилища!
func (r *itemRepository) GetItem(ctx context.Context, id int64) (*domain.Item, error) {
	r.RLock()
	defer r.RUnlock()
	log.Printf("Получаем item %d из хранилища! \n", id)

	return r.items[id], nil
}

// GetItem Получаем мапу items из хранилища!
func (r *itemRepository) GetItems(ctx context.Context) (*domain.Items, error) {
	r.Lock()
	defer r.Unlock()
	log.Println("Список!")

	return &r.items, nil
}
