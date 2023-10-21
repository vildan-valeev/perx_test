package domain

import (
	"perx/internal/transport/dto"
	"time"
)

type Item struct {
	ID            int     `json:"id"`             // ID элемента
	QueuePosition int     `json:"queue_position"` // номер в очереди
	Status        Status  `json:"status"`         // статус
	ElementsCount int     `json:"n"`              // количество элементов
	Delta         float64 `json:"d"`              // дельта между элементами полседовательности
	StartElement  float64 `json:"n1"`             // Стартовое значение
	TimeInterval  float64 `json:"I"`              // интервал в секундах
	TTL           float64 `json:"TTL"`            // время хранени результата в секуднах

	// TODO возможно в отдельную структуру...
	CurrentIteration int       `json:"current_iteration"` // текущая итерация
	ReceiptTime      time.Time `json:"receipt_time"`      // время поставки
	StartTime        time.Time `json:"start_time"`        // время запуска
	EndTime          time.Time `json:"end_time"`          // время окончания
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
func (i Items) ToDTO() []dto.ItemListDTO {
	var result []dto.ItemListDTO
	for _, item := range i {
		result = append(result, dto.ItemListDTO{
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
		})
	}

	return result
}
