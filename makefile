.PHONY: fmt lint run build watch

build:
	go build -o server main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run

fmt:
	goimports -w .
	go fmt ./...

lint:
	golangci-lint run

test:
	go test -v ./internal/utils
	go test -v ./internal/models

test_coverage:
	go test -cover -coverprofile=coverage/internal_utils.out ./internal/utils
	# go tool cover -html=coverage/internal_utils.out
	go test -cover -coverprofile=coverage/internal_models.out ./internal/models
	# go tool cover -html=coverage/internal_models.out