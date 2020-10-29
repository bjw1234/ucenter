package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonOK(c *gin.Context, msg string) {
	if msg == "" {
		msg = "ok"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": msg,
	})
}

func JsonErr(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": err.Error(),
	})
}
