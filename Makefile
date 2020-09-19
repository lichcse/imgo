dev-unit-test:
	go test ./...

dev-check-convention:
	find . -type d | xargs -L 1 golint

dev: dev-check-convention dev-unit-test

.DEFAULT_GOAL := dev
