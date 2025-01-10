package wrapper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func WrapSuccessResponse(c *gin.Context, data interface{}, message string) {
	response := SuccessResponse{
		StatusCode: http.StatusOK,
		Message:    message,
		Data:       data,
	}
	c.JSON(http.StatusOK, response)
}

func WrapErrorResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	response := SuccessResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
	c.JSON(statusCode, response)
}
