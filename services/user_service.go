package services

import (
	"api-go-gin/database"
	"api-go-gin/models"
	"api-go-gin/problem"
	"log"
)

func CreateUser(user *models.UserDTO) (*models.UserDTO, problem.ApplicationError) {

	if FindUserByUsername(user.Username) != nil {

		err := &problem.MessageErrorCode{
			Message:   "Username already exists.",
			ErrorCode: problem.User_Already_Exists_Error,
		}

		log.Println(err.Error())

		return nil, err
	}

	userModel := user.ToModel()

	err := database.DB.Create(userModel)

	if err.Error != nil {
		log.Println("Error creating user: ", err)
		return nil,
			&problem.MessageErrorCode{
				Message:   "Error creating user.",
				ErrorCode: problem.Database_Error,
			}
	}

	return userModel.ToDTO(), nil
}

func FindUserByUsername(username string) *models.UserDTO {
	user := models.User{
		Username: username,
	}

	database.DB.Where(&user).First(&user)

	if verifyIfUserExist(&user) {
		return user.ToDTO()
	}

	return nil
}

func verifyIfUserExist(user *models.User) bool {
	if user.ID == 0 {
		return false
	}
	return true
}
