package repository

import (
	"log"
	"sync"
	"time"

	"github.com/vildan-valeev/perx_test/internal/domain"
)

// Item - методы для работы с БД.
type LocalStorage interface {
	SetItem(item *domain.Item) error

	UpdateCurrentIteration(id int64, currentIteration int) error
	GetItem(id int64) (*domain.Item, error)
	GetItems() (*domain.Items, error)
	// Update Item
	SetStatus(id int64, status domain.Status) error
	SetStartTime(id int64, time time.Time) error
	SetEndTime(id int64, time time.Time) error
	DeleteItem(id int64) error
}

type localStorage struct {
	items domain.Items // задачи
	sync.RWMutex
}

func newStorage() *localStorage {
	return &localStorage{
		items: make(domain.Items),
	}
}

var instance *localStorage //nolint:gochecknoglobals
var once sync.Once         //nolint:gochecknoglobals

func GetLocalStorage() LocalStorage {
	once.Do(func() {
		instance = newStorage()
		log.Println("LocalStorage init...")
	})

	return instance
}

func (l *localStorage) UpdateCurrentIteration(id int64, currentIteration int) error {
	l.Lock()
	defer l.Unlock()
	log.Println("Обновление текущей итерации  item!")

	l.items[id].CurrentIteration = currentIteration

	return nil
}

// SetItem Добавляем item в хранилище!
func (l *localStorage) SetItem(item *domain.Item) error {
	l.Lock()
	defer l.Unlock()
	log.Printf("Добавляем item %d в хранилище! \n", item.ID)

	l.items[item.ID] = item

	return nil
}

// GetItem Получаем item из хранилища!
func (l *localStorage) GetItem(id int64) (*domain.Item, error) {
	l.RLock()
	defer l.RUnlock()
	log.Printf("Получаем item %d из хранилища! \n", id)

	return l.items[id], nil
}

// GetItems Получаем мапу items из хранилища.
func (l *localStorage) GetItems() (*domain.Items, error) {
	l.Lock()
	defer l.Unlock()
	log.Println("Получение списка!")

	return &l.items, nil
}

// SetStatus Обновляем поле статус.
func (l *localStorage) SetStatus(id int64, status domain.Status) error {
	l.Lock()
	defer l.Unlock()
	log.Printf("Обновляем поле Status для ID=%d \n", id)

	l.items[id].Status = status

	return nil
}

// SetStartTime Обновляем поле время начала обработки таски.
func (l *localStorage) SetStartTime(id int64, t time.Time) error {
	l.Lock()
	defer l.Unlock()
	log.Printf("Обновляем поле StartTime для ID=%d \n", id)

	l.items[id].StartTime = t

	return nil
}

// SetEndTime Обновляем поле время окончания обработки таски.
func (l *localStorage) SetEndTime(id int64, endTime time.Time) error {
	l.Lock()
	defer l.Unlock()
	log.Printf("Обновляем поле EndTime для ID=%d \n", id)

	l.items[id].EndTime = endTime

	return nil
}

// DeleteItem Удаление записи из мапы.
func (l *localStorage) DeleteItem(id int64) error {
	l.Lock()
	defer l.Unlock()
	log.Printf("Удаление элемента из списка ID=%d \n", id)

	delete(l.items, id)

	return nil
}
