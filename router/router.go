package routers

//
//import (
//	"net/http"
//
//	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/data/request"
//	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/service"
//	"github.com/gin-gonic/gin"
//)
//
//func RegisterAuthRoutes(router *gin.Engine, authService service.AuthService) {
//	router.POST("/register", func(c *gin.Context) {
//		var req request.RegisterUserRequest
//		if err := c.ShouldBindJSON(&req); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		resp, err := authService.Register(c, req)
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//			return
//		}
//		c.JSON(http.StatusOK, resp)
//	})
//
//	router.POST("/login", func(c *gin.Context) {
//		var req request.LoginRequest
//		if err := c.ShouldBindJSON(&req); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		resp, err := authService.Login(c, req)
//		if err != nil {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
//			return
//		}
//		c.JSON(http.StatusOK, resp)
//	})
//}
