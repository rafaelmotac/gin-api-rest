package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func returnStatusOkWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func returnStatusNotFoundWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": message,
	})
}

func returnStatusOkWithEntity(c *gin.Context, model any) {
	c.JSON(http.StatusOK, model)
}

func returnStatusCreatedWithEntity(c *gin.Context, model any) {
	c.JSON(http.StatusCreated, model)
}

func returnBadRequestWithValidationErrors(c *gin.Context, validationErrs []string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": validationErrs,
	})
}

func returnStatusBadRequestWithErr(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err,
	})
}

func returnStatusInternalServerErrorWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": message,
	})
}

func returnStatusConflictWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusConflict, gin.H{
		"error": message,
	})
}
