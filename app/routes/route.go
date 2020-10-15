package routes

import (
	v1 "imgo/app/routes/v1"
	"imgo/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var route = gin.Default()

// SetupRouter func setup app route
func SetupRouter(db *gorm.DB) *gin.Engine {
	identityModule(db)

	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return route
}

func identityModule(db *gorm.DB) {
	iV1 := route.Group("/identity/v1")
	{
		v1.UserRoute(iV1, db)
	}
}
