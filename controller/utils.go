package controller

import (
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/data/response"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"strings"
)

type ApiError struct {
	Code    int
	Message string
}

func (e *ApiError) Error() string {
	return e.Message
}

func HandleError(c *gin.Context, err error) {
	if apiErr, ok := err.(*ApiError); ok {
		JsonResponse(c, apiErr.Code, apiErr.Message, nil)
	} else {
		JsonResponse(c, http.StatusInternalServerError, "Internal Server Error", nil)
	}
}

func JsonResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, response.Response{
		Code:   status,
		Status: message,
		Data:   data,
	})
}

func GetClientIP(c *gin.Context) string {
	xForwardedFor := c.GetHeader("X-Forwarded-For")
	if xForwardedFor != "" {
		ips := strings.Split(xForwardedFor, ",")
		return strings.TrimSpace(ips[0]) // Возвращаем первый IP
	}

	xRealIP := c.GetHeader("X-Real-IP")
	if xRealIP != "" {
		return xRealIP
	}

	ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
	if err != nil {
		return c.Request.RemoteAddr
	}
	return ip
}
