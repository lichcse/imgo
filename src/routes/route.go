package routes

import (
	"imgo/docs"
	"imgo/src/database"
	v1 "imgo/src/routes/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var r = gin.Default()

// SetupRouter : router
func SetupRouter(db database.SQLDb) *gin.Engine {
	identityModule(db)

	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

// identity
func identityModule(db database.SQLDb) {
	iV1 := r.Group("/identity/v1")
	{
		v1.UserRoute(iV1, db)
	}
}
