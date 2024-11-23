package controller

import (
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/data/request"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/data/response"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService service.Service
}

func NewAuthController(authService *service.Service) *AuthController {
	return &AuthController{
		authService: *authService,
	}
}

// Register контроллер для регистрации пользователей
func (controller *AuthController) Register(c *gin.Context) {
	var userRequest request.RegisterUserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Invalid request payload",
			Data:   nil,
		})
		return
	}

	// Вызов Register метода из AuthService
	authResp, err := controller.authService.Register(c, userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "Registration successful",
		Data:   authResp,
	})
}

// Login контроллер для аутентификации пользователей
func (controller *AuthController) Login(c *gin.Context) {
	var loginRequest request.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Invalid request payload",
			Data:   nil,
		})
		return
	}

	// Вызов Login метода из AuthService
	authResp, err := controller.authService.Login(c, loginRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{
			Code:   http.StatusUnauthorized,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "Login successful",
		Data:   authResp,
	})
}

// RefreshToken контроллер для обновления access токена
func (controller *AuthController) RefreshToken(c *gin.Context) {
	var tokenRequest request.UpdateTokenRequest
	if err := c.ShouldBindJSON(&tokenRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Invalid request payload",
			Data:   nil,
		})
		return
	}

	// Вызов метода RefreshToken из AuthService
	newTokens, err := controller.authService.GetAccessToken(c, tokenRequest.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{
			Code:   http.StatusUnauthorized,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "Token refreshed successfully",
		Data:   newTokens,
	})
}
