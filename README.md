1. MODE:
   - https://github.com/gin-gonic/gin/blob/master/mode.go
   - export GIN_MODE=release
2. Config
   - config/development.yaml|etc...
3. Database
   - src/database/sql/*
4. Start
   - go run main.go dev|test|stg|prod
5. Test:
   - go test ./...
6. Swagger
   - swag init & restart server
   - http://localhost:8080/swagger/docs/index.html
   - ref: https://github.com/swaggo/swag
