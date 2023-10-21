package domain

import "time"

type Item struct {
	ID            int     `json:"id"`             // ID элемента
	QueuePosition int     `json:"queue_position"` // номер в очереди
	Status        Status  `json:"status"`         // статус
	ElementsCount int     `json:"n"`              // количество элементов
	Delta         float64 `json:"d"`              // дельта между элементами полседовательности
	StartElement  float64 `json:"n1"`             //Стартовое значение
	TimeInterval  float64 `json:"I"`              //интервал в секундах
	TTL           float64 `json:"TTL"`            // время хранени результата в секуднах

	// TODO возможно в отдельную структуру...
	CurrentIteration int       `json:"current_iteration"` // текущая итерация
	ReceiptTime      time.Time `json:"receipt_time"`      // время поставки
	StartTime        time.Time `json:"start_time"`        // время запуска
	EndTime          time.Time `json:"end_time"`          // время окончания
	err              error     // ошибка

}

type Items []*Item

// Status задачи
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

// TODO возможно раздельно..
type task struct {
	id          int
	idItem      int
	receiptTime time.Time
	startTime   time.Time
	endTime     time.Time
	err         error
}
