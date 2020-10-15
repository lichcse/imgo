package v1

import (
	"imgo/app/modules/identity/v1/handler/http"
	"imgo/app/modules/identity/v1/resource"
	"imgo/app/response"
	"imgo/app/utils"

	"imgo/app/modules/identity/v1/repository"
	"imgo/app/modules/identity/v1/service"
	"imgo/app/modules/identity/v1/validation"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var restResponse = response.NewRestResponse(utils.NewIMLanguage(resource.IdentityLang, resource.DefaultLang), resource.CodeMessageMapping)

// UserRoute : user route group
func UserRoute(r *gin.RouterGroup, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userHandler := http.NewUserHandler(service.NewUserService(userRepository), restResponse, validation.NewUserValidation())
	r.POST("/user", userHandler.Add)
	r.GET("/user/:id", userHandler.Detail)
	r.PUT("/user/:id", userHandler.Update)
	r.DELETE("/user/:id", userHandler.Delete)
}
