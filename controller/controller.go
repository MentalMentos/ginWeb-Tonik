package controller

import (
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/data/request"
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

//// Register контроллер для регистрации пользователей
//func (controller *AuthController) Register(c *gin.Context) {
//	var userRequest request.RegisterUserRequest
//	if err := c.ShouldBindJSON(&userRequest); err != nil {
//		c.JSON(http.StatusBadRequest, response.Response{
//			Code:   http.StatusBadRequest,
//			Status: "Invalid request payload",
//			Data:   nil,
//		})
//		return
//	}
//
//	// Вызов Register метода из AuthService
//	authResp, err := controller.authService.Register(c, userRequest)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, response.Response{
//			Code:   http.StatusInternalServerError,
//			Status: err.Error(),
//			Data:   nil,
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, response.Response{
//		Code:   http.StatusOK,
//		Status: "Registration successful",
//		Data:   authResp,
//	})
//}

func (controller *AuthController) Register(c *gin.Context) {
	var userRequest request.RegisterUserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		return
	}

	authResp, err := controller.authService.Register(c, userRequest)
	if err != nil {
		HandleError(c, err)
		return
	}

	JsonResponse(c, http.StatusOK, "Registration successful", authResp)
}

func (controller *AuthController) Login(c *gin.Context) {
	var userRequest request.LoginRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		return
	}

	authResp, err := controller.authService.Login(c, userRequest)
	if err != nil {
		HandleError(c, err)
		return
	}

	JsonResponse(c, http.StatusOK, "Auth successful", authResp)
}

func (controller *AuthController) UpdatePassword(c *gin.Context) {
	var userRequest request.UpdateUserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		return
	}

	authResp, err := controller.authService.UpdatePassword(c, userRequest)
	if err != nil {
		HandleError(c, err)
		return
	}

	JsonResponse(c, http.StatusOK, "Password updated successful", authResp)
}

func (controller *AuthController) RefreshToken(c *gin.Context) {
	var userRequest request.UpdateTokenRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		return
	}

	authResp, err := controller.authService.GetAccessToken(c, userRequest.RefreshToken)
	if err != nil {
		HandleError(c, err)
		return
	}

	JsonResponse(c, http.StatusOK, "Token refreshed successful", authResp)
}
