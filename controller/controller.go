package controller

import (
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/data/request"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/data/response"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Controller struct {
	userService service.UserService
}

func NewController(userService service.UserService) *Controller {
	return &Controller{
		userService: userService,
	}
}

func (controller *Controller) Create(c *gin.Context) {
	userRequest := request.CreateUserRequest{}
	err := c.ShouldBind(&userRequest) //извлекает данные из тела запроса
	if err != nil {
		log.Fatalf("create controller error", err)
	}
	controller.userService.Create(c, userRequest)
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

func (controller *Controller) Delete(c *gin.Context) {
	Id := c.Param("userId")
	id, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		log.Fatalf("delete controller error", err)
	}
	controller.userService.Delete(c, id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

func (controller *Controller) FindById(c *gin.Context) {
	Id := c.Param("userId")
	id, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		log.Fatalf("delete controller error", err)
	}
	controller.userService.FindById(c, id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

// FindAllTags 		godoc
// @Summary			Get All tags.
// @Description		Return list of tags.
// @Tags			tags
// @Success			200 {obejct} response.Response{}
// @Router			/tags [get]
func (controller *Controller) FindAll(c *gin.Context) {
	tagResponse := controller.userService.FindAll(c)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
