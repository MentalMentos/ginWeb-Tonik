package main

import (
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/config"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/controller"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/model"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/repository"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/service"
	routers "github.com/MentalMentos/ginWeb-Tonik/ginWeb/router"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"net/http"
)

func main() {

	router := gin.Default()
	//fc
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.User{})

	// Repository
	Repository := repository.NewRepo(db)

	// Service
	Service := service.New(Repository, validate)

	// Controller
	Controller := controller.NewController(Service)

	// Router
	routes := routers.NewRouter(Controller)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
