package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpStatus, code int, meessage string, data interface{}) {
	c.JSON(httpStatus, gin.H{
		"code":   code,
		"messge": meessage,
		"data":   data,
	})
}

func ResponseOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":   20001,
		"messge": "succeed",
		"data":   data,
	})
}

func ResponseCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"code":   20001,
		"messge": "create succeed",
		"data":   data,
	})
}

// TODO 业务 error 转化
func ResponseError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code":   500,
		"messge": "internal error",
		"error":  err,
	})
}
