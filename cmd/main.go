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
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home!")
	})

	db := config.DatabaseConnection()
	validate := validator.New()
	db.Table("users").AutoMigrate(&model.User{})

	authRepository := repository.NewRepo(db)
	authService := service.New(authRepository, validate)
	authController := controller.NewAuthController(authService)

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", authController.Register)             // Регистрация
		authRoutes.POST("/login", authController.Login)                   // Вход
		authRoutes.POST("/refresh", authController.RefreshToken)          // Обновление токена
		authRoutes.PUT("/update-password", authController.UpdatePassword) // Обновление пароля
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Main", "Failed to start server")
	}
}
