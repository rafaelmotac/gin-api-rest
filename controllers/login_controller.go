package controllers

import (
	"api-go-gin/database"
	"api-go-gin/models"
	"api-go-gin/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Login godoc
// @Summary Login
// @Schemes
// @Description A method to authentication and get a token
// @Tags Login
// @Accept json
// @Produce json
// @Param user body models.UserDTO true "User DTO"
// @Success 200 {object} string
// @Router /login [post]
func Login(c *gin.Context) {
	var userDTO models.UserDTO

	if bindingOk := VerifyAndBindUserDTO(c, &userDTO); !bindingOk {
		return
	}

	if noValidationErrors := ValidateUserDTO(c, userDTO); !noValidationErrors {
		return
	}

	var user models.User

	database.DB.Session(&gorm.Session{SkipHooks: true}).Model(&models.User{}).
		Where("username = ?", userDTO.Username).
		First(&user)

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDTO.Password)) != nil {
		returnStatusUnauthorizedWithMessage(c, "Invalid username or password.")
		return
	}

	token := util.GenerateToken(int(user.ID))

	c.JSON(200, gin.H{
		"token": token,
	})

}
