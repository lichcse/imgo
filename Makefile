dev-unit-test:
	go test ./...

dev-check-convention:
	find . -type d | xargs -L 1 golint

api-docs:
	swag init

start:
	go run main.go dev

dev: dev-check-convention dev-unit-test api-docs start

.DEFAULT_GOAL := dev
