# build stage
FROM golang:1.19 as builder

ENV GO111MODULE=on

WORKDIR /app

COPY ./app/go.mod .
COPY ./app/go.sum .
RUN go mod download

COPY app/ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /build/main ./cmd/app/main.go

FROM ubuntu:20.04

WORKDIR /

COPY --from=builder /build/main /build/main

EXPOSE 8000

ENTRYPOINT ["./build/main"]