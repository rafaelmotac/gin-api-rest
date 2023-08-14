package services

import (
	"api-go-gin/database"
	"api-go-gin/models"
)

func CreateStudent(student *models.Student) {
	database.DB.Create(&student)
}

func FindAllStudents(students *[]models.Student) {
	database.DB.Find(&students)
}

func FindStudentById(student *models.Student, id string) bool {
	database.DB.Find(&student, id)

	return verifyIfStudentExist(student)
}

func FindStudentByFiscalNumber(student *models.Student, fiscalNumber string) bool {
	database.DB.Where(&models.Student{FiscalNumber: fiscalNumber}).First(&student)

	return verifyIfStudentExist(student)
}

func UpdateStudentById(student *models.Student) {
	database.DB.Model(&student).Updates(student)
}

func EditStudent(student *models.Student) {
	database.DB.Model(&student).UpdateColumns(student)
}

func DeleteStudentById(student *models.Student, id string) bool {
	if exists := FindStudentById(student, id); !exists {
		return false
	}
	database.DB.Delete(&student)

	return true
}

func verifyIfStudentExist(student *models.Student) bool {
	if student.ID == 0 {
		return false
	}
	return true
}
