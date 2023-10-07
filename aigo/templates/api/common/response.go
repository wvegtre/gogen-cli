package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpStatus, code int, message string, data interface{}) {
	c.JSON(httpStatus, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

func ResponseOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    20000,
		"message": "succeed",
		"data":    data,
	})
}

func ResponseCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"code":   20100,
		"messge": "create succeed",
		"data":   data,
	})
}

func ResponseError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code":    500,
		"message": "internal error",
		"error":   err,
	})
}
