package domain

type Item struct {
	ID  int     // количество элементов
	n   int     // количество элементов
	d   float64 // дельта между элементами полседовательности
	n1  float64 //Стартовое значение
	i   float64 //интервал в секундах
	ttl float64 // время хранени результата в секуднах
}
