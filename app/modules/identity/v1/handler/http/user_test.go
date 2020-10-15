package http

import (
	"bytes"
	"encoding/json"
	"imgo/app/modules/identity/v1/service"
	"imgo/app/modules/identity/v1/validation"
	"imgo/app/response"
	dto "imgo/app/schema/identity/v1"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUserHandler_Add(t *testing.T) {
	response := new(response.IMResponseMock)
	userService := new(service.UserServiceMock)
	userValidation := new(validation.UserValidationMock)
	userHandler := NewUserHandler(userService, response, userValidation)

	ctx := gin.Context{}
	userAddRequest := &dto.UserAddRequest{
		FullName: "Test",
		Username: "test",
		Email:    "example@example.com",
		Password: "123456789",
	}
	body, _ := json.Marshal(userAddRequest)
	ctx.Request, _ = http.NewRequest("GET", "", bytes.NewBuffer(body))
	ctx.Request.Header.Set("Content-Type", "application/json; charset=utf-8")
	userDetailResponse := &dto.UserDetailResponse{}
	response.On("Out", &ctx, nil, userDetailResponse)
	userService.On("Add", userAddRequest).Return(userDetailResponse, nil)
	userValidation.On("Add", userAddRequest).Return(nil)

	userHandler.Add(&ctx)
}
