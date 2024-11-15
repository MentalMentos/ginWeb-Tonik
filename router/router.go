package routers

import (
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(controller *controller.Controller) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	baseRouter := router.Group("/api")

	// Маршруты для tags
	tagRouter := baseRouter.Group("/tags")
	tagRouter.GET("", controller.FindAll)
	tagRouter.GET("/:tagId", controller.FindById)
	tagRouter.POST("", controller.Create)
	tagRouter.PATCH("/:tagId", controller.Update)
	tagRouter.DELETE("/:tagId", controller.Delete)

	// Маршруты для users
	userRouter := baseRouter.Group("/users")
	userRouter.GET("", controller.FindAll)
	userRouter.POST("", controller.Create)

	// Добавляйте другие маршруты аналогично для других ресурсов

	return router
}
