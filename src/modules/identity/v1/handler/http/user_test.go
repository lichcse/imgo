package http

import (
	"bytes"
	"encoding/json"
	"imgo/src/common/identity/v1/dto"
	"imgo/src/modules/identity/v1/service"
	"imgo/src/response"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUserHandler_Add(t *testing.T) {
	response := new(response.IMResponseMock)
	userService := new(service.UserServiceMock)
	userHandler := NewUserHandler(userService, response)

	ctx := gin.Context{}
	user := dto.UserAdd{
		FullName: "Test",
		Username: "test",
		Email:    "example@example.com",
		Password: "123456789",
	}
	body, _ := json.Marshal(user)
	ctx.Request, _ = http.NewRequest("GET", "", bytes.NewBuffer(body))
	ctx.Request.Header.Set("Content-Type", "application/json; charset=utf-8")

	response.On("Out", &ctx, nil, nil)
	userService.On("Add", user).Return(nil)

	userHandler.Add(&ctx)
}