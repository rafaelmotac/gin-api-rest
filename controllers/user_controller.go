package controllers

import (
	"api-go-gin/models"
	"api-go-gin/problem"
	"api-go-gin/services"
	"api-go-gin/util"
	"github.com/gin-gonic/gin"
)

const ErrorCreatingUser = "Error creating user."

func CreateUser(c *gin.Context) {
	var userDTO models.UserDTO
	if bindingOk := VerifyAndBindUserDTO(c, &userDTO); !bindingOk {
		return
	}

	if noValidationErrors := ValidateUserDTO(c, userDTO); !noValidationErrors {
		return
	}

	userSaved, errorCode := services.CreateUser(&userDTO)

	if errorCode != nil && errorCode.Kind() == problem.User_Already_Exists_Error {
		returnStatusConflictWithMessage(c, "Username already exists.")
		return
	} else if errorCode != nil {
		returnStatusInternalServerErrorWithMessage(c, ErrorCreatingUser)
		return

	}

	returnStatusCreatedWithEntity(c, userSaved)

}

func GetUserByUsername(c *gin.Context) {
	username := c.Params.ByName("username")
	var userDTO *models.UserDTO

	userDTO = services.FindUserByUsername(username)

	if userDTO == nil {
		returnStatusNotFoundWithMessage(c, RecordNotFoundMessage)
		return
	}

	returnStatusOkWithEntity(c, userDTO)
}

func VerifyAndBindUserDTO(c *gin.Context, obj any) bool {
	err := util.VerifyAndBind(c, &obj)

	if err != nil {
		returnStatusBadRequestWithErr(c, err)
		return false
	}
	return true
}

func ValidateUserDTO(c *gin.Context, userDTO models.UserDTO) bool {
	validationErrs := util.ValidateModel(userDTO)

	if validationErrs != nil {
		returnBadRequestWithValidationErrors(c, validationErrs)
		return false
	}
	return true
}
