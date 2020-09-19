package v1

import (
	"imgo/src/modules/identity/v1/handler/http"
	"imgo/src/modules/identity/v1/resource"
	"imgo/src/utils"

	"imgo/src/modules/identity/v1/repository"
	"imgo/src/modules/identity/v1/service"

	"imgo/src/response"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var res = response.NewRestResponse(utils.NewIMLanguage(resource.IdentityLang, resource.DefaultLang), resource.CodeMessageMapping)

// UserRoute : user route
func UserRoute(r *gin.RouterGroup, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userHandler := http.NewUserHandler(service.NewUserService(userRepository), res)
	r.POST("/user", userHandler.Add)
	r.GET("/user/:id", userHandler.Detail)
	r.PUT("/user/:id", userHandler.Update)
	r.DELETE("/user/:id", userHandler.Delete)
}
