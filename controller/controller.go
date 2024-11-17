package controller

import (
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/data/request"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/data/response"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Controller struct {
	userService service.Service
}

func NewController(userService service.Service) *Controller {
	return &Controller{
		userService: userService,
	}
}

func (controller *Controller) Create(c *gin.Context) {
	userRequest := request.RegisterUserRequest{}
	err := c.ShouldBind(&userRequest) //извлекает данные из тела запроса
	if err != nil {
		log.Fatalf("create controller error", err)
	}
	controller.userService.Register(c, userRequest)
	response := response.Response{
		http.StatusOK,
		"Ok",
		nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func (controller *Controller) Update(c *gin.Context) {
	userRequest := request.UpdateUserRequest{}
	err := c.ShouldBind(&userRequest) //извлекает данные из тела запроса
	if err != nil {
		log.Fatalf("update controller error", err)
	}
	controller.userService.Update(c, userRequest)
	response := response.Response{
		http.StatusOK,
		"Ok",
		nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}
