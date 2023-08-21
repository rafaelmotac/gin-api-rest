package routes

import (
	"api-go-gin/controllers"
	docs "api-go-gin/docs"
	"github.com/gin-contrib/cors"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	}))
	r.Static("/static", "./static")

	r.GET("/app/index", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.LoadHTMLGlob("templates/*")
	r.GET("/index", controllers.Index)
	r.GET("/", controllers.Greeting)
	r.GET("/students", controllers.GetAllStudents)
	r.GET("/students/:id", controllers.GetStudent)
	r.GET("/students/fiscal_number/:fiscalNumber", controllers.GetStudentByFiscalNumber)
	r.POST("/students", controllers.CreateStudent)
	r.PUT("/students/:id", controllers.UpdateStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:username", controllers.GetUserByUsername)
	r.Run(":9000")
}
