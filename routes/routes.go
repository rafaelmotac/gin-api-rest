package routes

import (
	"api-go-gin/controllers"
	docs "api-go-gin/docs"
	"api-go-gin/middleware"
	"api-go-gin/properties"
	"fmt"
	"github.com/gin-contrib/cors"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"

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

	r.POST("/login", controllers.Login)
	r.GET("/", controllers.Greeting)

	private := r.Group("/")
	private.Use(middleware.JWTAuthMiddleware())

	private.GET("/students", controllers.GetAllStudents)
	private.GET("/students/:id", controllers.GetStudent)
	private.GET("/students/fiscal_number/:fiscalNumber", controllers.GetStudentByFiscalNumber)
	private.POST("/students", controllers.CreateStudent)
	private.PUT("/students/:id", controllers.UpdateStudent)
	private.PATCH("/students/:id", controllers.EditStudent)
	private.DELETE("/students/:id", controllers.DeleteStudent)
	private.POST("/users", controllers.CreateUser)
	private.GET("/users/:username", controllers.GetUserByUsername)
	err := r.Run(fmt.Sprintf(":%d", properties.Properties.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}
