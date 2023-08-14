package controllers

import (
	"api-go-gin/models"
	"api-go-gin/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

const RecordNotFoundMessage = "Record not found."

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Home Page",
		"message": "Hello World!",
	})
}

// GetAllStudents godoc
// @Summary Get All Students
// @Schemes
// @Description A method to get all students
// @Tags Students
// @Accept json
// @Produce json
// @Success 200 {object} models.Student
// @Router /students [get]
func GetAllStudents(c *gin.Context) {
	var students []models.Student
	services.FindAllStudents(&students)

	returnStatusOkWithEntity(c, students)
}

// GetStudent godoc
// @Summary Get Student by ID
// @Schemes
// @Description A method to get a student by ID
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} models.Student
// @Router /students/{id} [get]
func GetStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student

	if found := services.FindStudentById(&student, id); !found {
		returnStatusNotFoundWithMessage(c, RecordNotFoundMessage)
		return
	}

	returnStatusOkWithEntity(c, student)
}

// GetStudentByFiscalNumber godoc
// @Summary Get Student by Fiscal Number
// @Schemes
// @Description A method to get a student by Fiscal Number
// @Tags Students
// @Accept json
// @Produce json
// @Param fiscalNumber path string true "Student Fiscal Number"
// @Success 200 {object} models.Student
// @Router /students/fiscal_number/{fiscalNumber} [get]
func GetStudentByFiscalNumber(c *gin.Context) {
	var student models.Student
	fiscalNumber := c.Params.ByName("fiscalNumber")

	if found := services.FindStudentByFiscalNumber(&student, fiscalNumber); !found {
		returnStatusNotFoundWithMessage(c, RecordNotFoundMessage)
		return
	}

	returnStatusOkWithEntity(c, student)

}

// CreateStudent godoc
// @Summary Create Student
// @Schemes
// @Description A method to create a student
// @Tags Students
// @Accept json
// @Produce json
// @Param student body models.Student true "Student"
// @Success 201 {object} models.Student
// @Router /students [post]
func CreateStudent(c *gin.Context) {
	var student models.Student
	if bindingOk := verifyAndBind(c, &student); !bindingOk {
		return
	}

	if noValidationErrors := validateStudent(c, student); !noValidationErrors {
		return
	}

	services.CreateStudent(&student)
	returnStatusCreatedWithEntity(c, student)
}

// UpdateStudent godoc
// @Summary Update Student
// @Schemes
// @Description A method to update a student
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Param student body models.Student true "Student"
// @Success 200 {object} models.Student
// @Router /students/{id} [put]
func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	if exists := services.FindStudentById(&student, id); !exists {
		returnStatusNotFoundWithMessage(c, RecordNotFoundMessage)
		return
	}

	if bindingOk := verifyAndBind(c, &student); !bindingOk {
		return
	}

	if noValidationErrors := validateStudent(c, student); !noValidationErrors {
		return
	}

	services.UpdateStudentById(&student)

	returnStatusOkWithEntity(c, student)
}

// EditStudent godoc
// @Summary Edit Student
// @Schemes
// @Description A method to edit a student
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Param student body models.Student true "Student"
// @Success 200 {object} models.Student
// @Router /students/{id} [patch]
func EditStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	if exists := services.FindStudentById(&student, id); !exists {
		returnStatusNotFoundWithMessage(c, RecordNotFoundMessage)
		return
	}

	if bindingOk := verifyAndBind(c, &student); !bindingOk {
		return
	}

	if noValidationErrors := validateStudent(c, student); !noValidationErrors {
		return
	}

	services.EditStudent(&student)

	returnStatusOkWithEntity(c, student)
}

// DeleteStudent godoc
// @Summary Delete Student
// @Schemes
// @Description A method to delete a student
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Success 200
// @Router /students/{id} [delete]
func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	if exists := services.DeleteStudentById(&student, id); !exists {
		returnStatusNotFoundWithMessage(c, RecordNotFoundMessage)
		return
	}

	returnStatusOkWithMessage(c, "Student deleted successfully")
}

func validateStudent(c *gin.Context, student models.Student) bool {
	validationErrs := services.ValidateModel(student)

	if validationErrs != nil {
		returnBadRequestWithValidationErrors(c, validationErrs)
		return false
	}
	return true
}

func verifyAndBind(c *gin.Context, student *models.Student) bool {
	err := c.ShouldBindJSON(student)

	if err != nil {
		returnStatusBadRequestWithErr(c, err)
		return false
	}
	return true
}
