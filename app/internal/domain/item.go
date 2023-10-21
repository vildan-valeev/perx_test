package domain

import (
	"time"

	"github.com/vildan-valeev/perx_test/internal/transport/dto"
)

type Item struct {
	ID            int64   // ID элемента
	QueuePosition int     // номер в очереди
	Status        Status  // статус
	ElementsCount int     // количество элементов
	Delta         float64 // дельта между элементами полседовательности
	StartElement  float64 // Стартовое значение
	TimeInterval  float64 // интервал в секундах
	TTL           float64 // время хранени результата в секуднах

	// TODO возможно в отдельную структуру...
	CurrentIteration int       // текущая итерация
	ReceiptTime      time.Time // время поставки
	StartTime        time.Time // время запуска
	EndTime          time.Time // время окончания
	err              error     // ошибка
}

type Items []*Item

// Status задачи.
type Status int

const (
	StatusUnknown   Status = iota // Зарезервированный код (не используемое значение, статус неизвестно)
	StatusProcessed               // В процессе
	StatusInQueue                 // В очереди
	StatusDone                    // Завершен
	StatusError                   // Ошибка
)

func (s Status) String() string {
	switch s {
	case 0:
		return "Unknown"
	case 1:
		return "В процессе"
	case 2:
		return "В очереди"
	case 3:
		return "Завершен"
	case 4:
		return "Ошибка"
	default:
		return ""
	}
}

// ToDTO TODO перенесте в dto, сделать передачу аргументом а не методом.
func (i Items) ToDTO() dto.ItemsDTO {
	result := dto.ItemsDTO{}

	for _, item := range i {
		result = append(result, dto.ItemDTO{
			ID:               item.ID,
			QueuePosition:    item.QueuePosition,
			Status:           item.Status.String(),
			ElementsCount:    item.ElementsCount,
			Delta:            item.Delta,
			StartElement:     item.StartElement,
			TimeInterval:     item.TimeInterval,
			TTL:              item.TTL,
			CurrentIteration: item.CurrentIteration,
			ReceiptTime:      item.ReceiptTime.UnixMilli(),
			StartTime:        item.StartTime.UnixMilli(),
			EndTime:          item.EndTime.UnixMilli(),
			Err:              item.err.Error(),
		})
	}

	return result
}
