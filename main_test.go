package main

import (
	"api-go-gin/config"
	"api-go-gin/controllers"
	"api-go-gin/database"
	"api-go-gin/models"
	"api-go-gin/properties"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var IdTestStudent int

func SetupTestingRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func TestVerifyStatusCode(t *testing.T) {
	properties.InitProperties()
	r := SetupTestingRoutes()
	r.GET("/", controllers.Greeting)

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")

	mockResponse := `{"message":"Hello World!"}`
	assert.Equal(t, mockResponse, resp.Body.String(), "Response body is not expected")
}

func TestGetAllStudents(t *testing.T) {
	properties.InitProperties()
	database.DbConnect()
	createMockStudent()
	defer deleteMockStudent()
	r := SetupTestingRoutes()
	r.GET("/students", controllers.GetAllStudents)

	req, _ := http.NewRequest(http.MethodGet, "/students", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var studentsResponse []models.Student
	json.Unmarshal(resp.Body.Bytes(), &studentsResponse)

	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
	assert.Contains(t, resp.Body.String(), "Test Student", "Response body is not expected")

}

func TestGetStudentByFiscalNumber(t *testing.T) {
	properties.InitProperties()
	database.DbConnect()
	createMockStudent()
	defer deleteMockStudent()
	r := SetupTestingRoutes()
	r.GET("/students/fiscal_number/:fiscalNumber", controllers.GetStudentByFiscalNumber)

	req, _ := http.NewRequest(http.MethodGet, "/students/fiscal_number/98732165478", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var studentResponse models.Student
	json.Unmarshal(resp.Body.Bytes(), &studentResponse)

	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
	assert.Equal(t, "Test Student", studentResponse.Name, "Student name is not expected")
	assert.Equal(t, "98732165478", studentResponse.FiscalNumber, "Student fiscal number is not expected")
	assert.Equal(t, "987654321", studentResponse.IdentificationNumber, "Student identification number is not expected")

}

func TestGetStudentById(t *testing.T) {
	properties.InitProperties()
	database.DbConnect()
	createMockStudent()
	defer deleteMockStudent()
	r := SetupTestingRoutes()
	r.GET("/students/:id", controllers.GetStudent)

	path := "/students/" + strconv.Itoa(IdTestStudent)
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var studentResponse models.Student
	json.Unmarshal(resp.Body.Bytes(), &studentResponse)

	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
	assert.Equal(t, "Test Student", studentResponse.Name, "Student name is not expected")
	assert.Equal(t, "98732165478", studentResponse.FiscalNumber, "Student fiscal number is not expected")
	assert.Equal(t, "987654321", studentResponse.IdentificationNumber, "Student identification number is not expected")

}

func TestUpdateStudentById(t *testing.T) {
	properties.InitProperties()
	config.InitValidator()
	database.DbConnect()
	createMockStudent()
	defer deleteMockStudent()
	r := SetupTestingRoutes()
	r.PUT("/students/:id", controllers.UpdateStudent)

	student := models.Student{
		Name:                 "Test Student Updated",
		FiscalNumber:         "12345678901",
		IdentificationNumber: "123456789",
	}

	studentJson, _ := json.Marshal(student)
	path := "/students/" + strconv.Itoa(IdTestStudent)
	req, _ := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(studentJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var studentResponse models.Student
	err := json.Unmarshal(resp.Body.Bytes(), &studentResponse)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
	assert.Equal(t, student.Name, studentResponse.Name, "Student name is not expected")
	assert.Equal(t, student.FiscalNumber, studentResponse.FiscalNumber, "Student fiscal number is not expected")
	assert.Equal(t, student.IdentificationNumber, studentResponse.IdentificationNumber, "Student identification number is not expected")
}

func TestEditStudentById(t *testing.T) {
	properties.InitProperties()
	config.InitValidator()
	database.DbConnect()
	createMockStudent()
	defer deleteMockStudent()
	r := SetupTestingRoutes()
	r.PATCH("/students/:id", controllers.EditStudent)

	student := models.Student{
		Name: "Test Student Updated",
	}

	studentJson, _ := json.Marshal(student)
	path := "/students/" + strconv.Itoa(IdTestStudent)
	req, _ := http.NewRequest(http.MethodPatch, path, bytes.NewBuffer(studentJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var studentResponse models.Student
	err := json.Unmarshal(resp.Body.Bytes(), &studentResponse)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
	assert.Equal(t, student.Name, studentResponse.Name, "Student name is not expected")
}

func TestCreateStudent(t *testing.T) {
	properties.InitProperties()
	config.InitValidator()
	database.DbConnect()
	r := SetupTestingRoutes()
	r.POST("/students", controllers.CreateStudent)

	student := models.Student{
		Name:                 "Test Student",
		FiscalNumber:         "98732165478",
		IdentificationNumber: "987654321",
	}

	studentJson, _ := json.Marshal(student)
	req, _ := http.NewRequest(http.MethodPost, "/students", bytes.NewBuffer(studentJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var studentResponse models.Student
	err := json.Unmarshal(resp.Body.Bytes(), &studentResponse)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	IdTestStudent = int(studentResponse.ID)
	defer deleteMockStudent()

	assert.Equal(t, http.StatusCreated, resp.Code, "Created response is expected")
	assert.Equal(t, student.Name, studentResponse.Name, "Student name is not expected")
	assert.Equal(t, student.FiscalNumber, studentResponse.FiscalNumber, "Student fiscal number is not expected")
	assert.Equal(t, student.IdentificationNumber, studentResponse.IdentificationNumber, "Student identification number is not expected")

}

func TestDeleteStudent(t *testing.T) {
	properties.InitProperties()
	database.DbConnect()
	createMockStudent()
	defer deleteMockStudent()
	r := SetupTestingRoutes()
	r.DELETE("/students/:id", controllers.DeleteStudent)

	path := "/students/" + strconv.Itoa(IdTestStudent)
	req, _ := http.NewRequest(http.MethodDelete, path, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var messageResponse string

	expectedMessage := `{"message":"Student deleted successfully"}`

	messageResponse = resp.Body.String()

	assert.Equal(t, http.StatusOK, resp.Code, "No content response is expected")
	assert.Equal(t, expectedMessage, messageResponse, "Response message is not expected")

}

func createMockStudent() {
	mockStudent := models.Student{
		Name:                 "Test Student",
		FiscalNumber:         "98732165478",
		IdentificationNumber: "987654321",
	}
	database.DB.Create(&mockStudent)
	IdTestStudent = int(mockStudent.ID)
}

func deleteMockStudent() {
	database.DB.Unscoped().Delete(&models.Student{}, IdTestStudent)
}
