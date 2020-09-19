dev-unit-test:
	go test ./...

dev-check-convention:
	find . -type d | xargs -L 1 golint

swag:
	swag init

start:
	go run main.go dev

dev: dev-check-convention dev-unit-test swag start

.DEFAULT_GOAL := dev
