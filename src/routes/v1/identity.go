package v1

import (
	"imgo/src/modules/identity/v1/handler/http"
	"imgo/src/modules/identity/v1/resource"
	"imgo/src/utils"

	"imgo/src/modules/identity/v1/repository"
	"imgo/src/modules/identity/v1/service"

	"imgo/src/database"
	"imgo/src/response"

	"github.com/gin-gonic/gin"
)

var res = response.NewRestResponse(utils.NewIMLanguage(resource.IdentityLang, resource.DefaultLang), resource.CodeMessageMapping)

// UserRoute : user route
func UserRoute(r *gin.RouterGroup, db database.SQLDb) {
	userRepository := repository.NewUserRepository(db)
	userHandler := http.NewUserHandler(service.NewUserService(userRepository), res)
	r.POST("/user", userHandler.Add)
	r.GET("/user/:user_id", userHandler.Detail)
	r.PUT("/user/:user_id", userHandler.Update)
	r.DELETE("/user/:user_id", userHandler.Delete)
}
