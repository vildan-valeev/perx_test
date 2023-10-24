package repository

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/vildan-valeev/perx_test/internal/domain"
)

type itemRepository struct {
	items    domain.Items // задачи
	resultCh <-chan domain.Item
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

func (r *itemRepository) GetResultCan() <-chan domain.Item {
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

// GetItems Получаем мапу items из хранилища.
func (r *itemRepository) GetItems(ctx context.Context) (*domain.Items, error) {
	r.Lock()
	defer r.Unlock()
	log.Println("Список!")

	return &r.items, nil
}

// SetStatus Обновляем поле статус.
func (r *itemRepository) SetStatus(id int64, status domain.Status) error {
	r.Lock()
	defer r.Unlock()
	log.Println("Обновляем поле Status")

	return nil
}

// SetStartTime Обновляем поле время начала обработки таски.
func (r *itemRepository) SetStartTime(id int64, time time.Time) error {
	r.Lock()
	defer r.Unlock()
	log.Println("Обновляем поле StartTime")

	return nil
}

// SetEndTime Обновляем поле время окончания обработки таски.
func (r *itemRepository) SetEndTime(id int64, endTime time.Time) error {
	r.Lock()
	defer r.Unlock()
	log.Println("Обновляем поле EndTime")

	r.items[id].EndTime = endTime

	return nil
}

// DeleteItem Удаление записи из мапы.
func (r *itemRepository) DeleteItem(id int64) error {
	r.Lock()
	defer r.Unlock()
	log.Println("Удаление элемента из списка!")

	delete(r.items, id)

	return nil
}
