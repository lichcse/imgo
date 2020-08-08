package routes

import (
	"imgo/docs"
	v1 "imgo/src/routes/v1"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var r = gin.Default()

// SetupRouter : router
func SetupRouter(db *gorm.DB) *gin.Engine {
	identityModule(db)

	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

// identity
func identityModule(db *gorm.DB) {
	iV1 := r.Group("/identity/v1")
	{
		v1.UserRoute(iV1, db)
	}
}
