package service

import (
	"context"
	"crypto/rand"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/vildan-valeev/perx_test/internal/domain"
	"github.com/vildan-valeev/perx_test/internal/repository"
	"github.com/vildan-valeev/perx_test/internal/transport/dto"
	"github.com/vildan-valeev/perx_test/pkg/pool"
)

// ItemService UseCase - бизнес логика.
type ItemService struct {
	wp   *pool.Pool // Worker Pool
	repo repository.Item
}

func NewItemService(repo repository.Item, wp *pool.Pool) *ItemService {
	return &ItemService{
		repo: repo,
		wp:   wp,
	}
}

// AddItemToQueueService Добавление задачи в очередь.
func (s ItemService) AddItemToQueueService(ctx context.Context, addItem *dto.ItemToQueueDTO) error {
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
		ReceiptTime:   time.Now(), // выставляем время получения item в обработку
	}

	// todo: validation business logic Item fields

	// отправляем на обработку в очередь worker pool
	task := pool.NewTask(s.progression, &ArgsProgression{Item: i})
	s.wp.AddTask(task) // TODO: add errors

	i.Status = domain.StatusInQueue
	// сохраняем в память(оправляем в хранилище/бд)
	if err = s.repo.SetItem(ctx, &i); err != nil {
		return err
	}

	return nil
}

// ListItemService Получение списка.
func (s ItemService) ListItemService(ctx context.Context) (*domain.Items, error) {
	return s.repo.GetItems(ctx)
}

// ArgsProgression Аргументы для проброса в таск и вызова функции progression из таски в воркер пуле.
type ArgsProgression struct {
	Item domain.Item // Отдаем копию в обработку(по ссылке только в/из хранилища)
}

func (s ItemService) progression(arguments interface{}) error {
	// TODO: сделать проверки интерфейса на наличие аргументов
	args := arguments.(*ArgsProgression)
	id := args.Item.ID
	n := args.Item.ElementsCount
	d := args.Item.Delta
	a1 := args.Item.StartElement
	I := args.Item.TimeInterval

	// TODO: РЕЗУЛЬТАТЫ ЧЕРЕЗ КАНАЛ ПРОТАСКИВАЕМ ИЛИ ТАК ДЕРГАЕМ МЕТОДЫ ?????

	log.Printf("Start iteration ID=%d, start=%f, delta=%f, n=%d I=%f\n", id, a1, d, n, I)

	if err := s.repo.SetStatus(id, domain.StatusProcessed); err != nil {
		return err
	}

	if err := s.repo.SetStartTime(id, time.Now()); err != nil {
		return err
	}

	for i := 1; i < n+1; i++ {
		res := a1 + (d * (float64(i) - 1))
		log.Printf("Iteration... ID=%d, curIter=%d, res=%f \n", id, i, res)
		// обновить текущее значение итерации Item в хранилище
		if err := s.repo.UpdateCurrentIteration(id, i); err != nil {
			return err
		}
		// интервал в секундах между итерациями
		time.Sleep(time.Duration(I) * time.Second)
	}

	if err := s.repo.SetStatus(id, domain.StatusDone); err != nil {
		return err
	}

	if err := s.repo.SetEndTime(id, time.Now()); err != nil {
		return err
	}

	return nil
}

//nolint:gochecknoglobals
var doOnce sync.Once

func (s ItemService) ProcessingResults(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Minute)

	doOnce.Do(func() {
		log.Println("Запуск воркера обработки хранилища")
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					err := s.cleaning(ctx)
					if err != nil {
						log.Println(err.Error())
					}
				}
			}
		}()
	})

	log.Println("Обработка результатов уже запущена")
}

func (s ItemService) cleaning(ctx context.Context) error {
	log.Printf("Удаление старых данных...")

	items, err := s.repo.GetItems(ctx)
	if err != nil {
		return err
	}

	for k, v := range *items {
		if v.EndTime.After(time.Now().Add(-time.Duration(v.TTL) * time.Second)) {
			err = s.repo.DeleteItem(k)
			if err != nil {
				return err
			}
		}
	}

	log.Println("Обработкa завершена...")

	return nil
}
