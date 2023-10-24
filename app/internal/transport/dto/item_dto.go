package dto

type ItemToQueueDTO struct {
	N   int     `json:"n"`   // Количество элементов
	D   float64 `json:"d"`   // Дельта между элементами полседовательности
	N1  float64 `json:"n1"`  // Стартовое значение
	I   float64 `json:"I"`   // Интервал в секундах
	TTL float64 `json:"TTL"` // время хранени результата в секуднах
}

type ItemDTO struct {
	ID               int64   `json:"id"`                // ID элемента
	QueuePosition    int     `json:"queue_position"`    // Номер в очереди
	Status           string  `json:"status"`            // Статус
	ElementsCount    int     `json:"n"`                 // Количество элементов
	Delta            float64 `json:"d"`                 // Дельта между элементами полседовательности
	StartElement     float64 `json:"n1"`                // Стартовое значение
	TimeInterval     float64 `json:"I"`                 // Интервал в секундах
	TTL              float64 `json:"TTL"`               // время хранени результата в секуднах
	CurrentIteration int     `json:"current_iteration"` // текущая итерация
	ReceiptTime      int64   `json:"receipt_time"`      // Время поставки(unix)
	StartTime        int64   `json:"start_time"`        // Время запуска (unix)
	EndTime          int64   `json:"end_time"`          // Время окончания (unix)
}

type ItemsDTO []ItemDTO
