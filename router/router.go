package routers

import (
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(Controller *controller.Controller) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.GET("", Controller.FindAll)
	tagsRouter.GET("/:tagId", Controller.FindById)
	tagsRouter.POST("", Controller.Create)
	tagsRouter.PATCH("/:tagId", Controller.Update)
	tagsRouter.DELETE("/:tagId", Controller.Delete)

	return router
}
