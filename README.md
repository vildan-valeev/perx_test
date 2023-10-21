# Тестовое задание на Backend-разработчика на Go
Написать сервис, который считает арифметическую прогрессию в очереди.
Задачи поступают в очередь, из очереди поступют на выполнение, выполняются
до получения результата, после чего из очереди выбирается следующая задача.
Параллельно может выполняться N задач.
Нужно создать github репозиторий и в нем разместить приложение. При запуске
стартует HTTP-сервер, у сервера есть два endpointa:
1. Поставить задачу в очередь. Параметры:
    * n - количество элементов (целочисленное)
    * d - дельта между элементами последовательности (вещественное)
    * n1 - Стартовое значение (вещественное)
    * I - интервал в секундах между итерациями (вещественное)
    * TTL - время хранения результата в секундах (вещественное)
2. Получить отсортированный список задач и статусы выполнения этих задач. Поля результата для каждой задачи:
    * Номер в очереди (целочисленное)
    * Статус: В процессе/В очереди/Завершена 
    * n 
    * d 
    * n1 
    * I 
    * TTL 
    * Текущая итерация 
    * Время постановки задачи 
    * Время старта задачи 
    * Время окончания задачи (в случае если задача завершена)
### Требования к реализации
* Отработанные задачи стирать после завершения TTL;
* В качестве хранения данных использовать память;
* Вычисление текущего значения должно высчитываться от предыдущего
значения по факту наступления времени (по интервалу), а не по формуле
разницы времени и количеству итераций;
* Параллельно может выполняться только N задач. Количество N передается
через параметры командной строки сервиса.
### Оценивается (кроме правильности решения):
- качество и читаемость кода
- структура кода
- тесты и проверяемость результата
- готовность к production (правильный страт и остановка), параметры запуска   

Приветствуется использование только стандартной библиотеки.

---
## Prerequisites
Golang  
Docker  
Golangci-lint (optional)  
Make  
brains and hands
---
## Start
Instruction
```shell
make help
```
Start in local
```shell
make up_local
```
Start in docker
```shell
make up
```
Lint
```shell
make lint
```
Test
```shell
make test
```
Data race detector:
```shell
make test_race
```
---
## Architecture

Слой transport - принимаем данные, запросы, проверка входящих значение, укладка в DTO, формирование ответов   
Слой service - запуск задачи в worker pools, прием и запись результатов в память, выдача результатов в транспорт  
слой repo - работа с внешними сервисами(бд, память), работа с worker_pool  

старт worker pool вместе с приложением из main
реализация worker pool из pkg  
итерации с прогрессией в service(или pkg)

---
### TODO
1. обработка ошибок
2. возможно следует переименовать в "а1" поле n1 (startElement)
3. заменить в импортах "perx" на "github.com/..../perx_test"
