
.PHONY: help
help:
	@echo "-Комп перезагружал?\n-Да!\n-Клаву протирал?\n-Да!\n-Тогда не знаю в чем проблема..."

.PHONY: up
up:
	cd app && docker-compose up --build


.PHONY: up_local
up_local:
	go build -o /build/main ./cmd/app/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /build/main ./cmd/app/main.go

.PHONY: lint
lint:
	cd app && golangci-lint run --config ./build/.golangci.yml ./...

.PHONY: test
test:
	cd app && go test ./...

.PHONY: test_race
test_race:
	cd app && go test -race -short ./...
