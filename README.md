1. MODE:
   - https://github.com/gin-gonic/gin/blob/master/mode.go
   - export GIN_MODE=release
2. Config
   - config/development.yaml|etc...
2. Start
   - go run main.go dev|test|stg|prod
3. Test:
   - cd test/
   - go test ./...
4. Swagger
   - swag init & restart server
   - http://localhost:8080/swagger/docs/index.html
   - ref: https://github.com/swaggo/swag
