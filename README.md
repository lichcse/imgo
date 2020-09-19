# Simple project template
## Overview
This is a simple project template with Go.
## Quick start
* Download or clone template project
* Config environment file `config/dev.yaml`
* Start server `go run main.go dev` or `make`
* Using swagger test api `http://localhost:{{PORT}}/swagger/docs/index.html`
## Common command
* Start server: `go run main.go dev|test|stg|prd`
* Unit test: `go test ./...` or `make dev-unit-test`
* Check  convention: `find . -type d | xargs -L 1 golint` or `make dev-check-convention`
