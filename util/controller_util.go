package util

import (
	"github.com/gin-gonic/gin"
)

func VerifyAndBind(c *gin.Context, model *any) error {
	return c.ShouldBindJSON(&model)
}
