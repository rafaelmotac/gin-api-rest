package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Greeting godoc
// @Summary Greeting
// @Schemes
// @Description A method to get a greeting
// @Tags Greeting
// @Accept json
// @Produce json
// @Success 200 {object} string
func Greeting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}
