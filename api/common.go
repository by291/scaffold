package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// For some common response apis

func Success(c *gin.Context, code int) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  http.StatusText(code),
	})
}

func SuccessWithData(c *gin.Context, code int, data any) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  http.StatusText(code),
		"data": data,
	})
}

func Fail(c *gin.Context, code int) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  http.StatusText(code),
	})
}

func FailWithError(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  err.Error(),
	})
}
