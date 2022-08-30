package utils

import "github.com/gin-gonic/gin"

func ConverterJson(c *gin.Context, status int, message string) {
	c.JSON(status, message)
}
