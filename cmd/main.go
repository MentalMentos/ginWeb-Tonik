package main

import (
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/config"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/controller"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/model"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/repository"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/service"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"log"
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
	service := service.New(Repository, validate)

	// Controller
	controller := controller.NewAuthController(service)

	router.POST("/tasks", func(c *gin.Context) { controller.Register(c) })
	//	log.Info("Main", "Starting server on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Main", "Failed to start server")
	}

}
