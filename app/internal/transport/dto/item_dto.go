package dto

type ItemToQueueDTO struct {
	N   int     `json:"n"`   // количество элементов
	D   float64 `json:"d"`   // дельта между элементами полседовательности
	N1  float64 `json:"n1"`  // Стартовое значение
	I   float64 `json:"I"`   // интервал в секундах
	TTL float64 `json:"TTL"` // время хранени результата в секуднах
}

type ItemListDTO struct {
	ID               int     `json:"id"`                // ID элемента
	QueuePosition    int     `json:"queue_position"`    // номер в очереди
	Status           string  `json:"status"`            // статус
	ElementsCount    int     `json:"n"`                 // количество элементов
	Delta            float64 `json:"d"`                 // дельта между элементами полседовательности
	StartElement     float64 `json:"n1"`                //Стартовое значение
	TimeInterval     float64 `json:"I"`                 //интервал в секундах
	TTL              float64 `json:"TTL"`               // время хранени результата в секуднах
	CurrentIteration int     `json:"current_iteration"` // текущая итерация
	ReceiptTime      int64   `json:"receipt_time"`      // время поставки(unix)
	StartTime        int64   `json:"start_time"`        // время запуска (unix)
	EndTime          int64   `json:"end_time"`          // время окончания (unix)

	// 400й код, с телом - бизнесовые ошибки!!! все остальное - сеть
	Err error `json:"error"` // ошибка

}
