package main

import (
	"gowebservice/config"
	"gowebservice/controllers"

	"github.com/gin-gonic/gin"
)

var err error

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	r.GET("/students", controllers.GetStudents)
	r.GET("/student/:id", controllers.GetStudentById)
	r.POST("/student", controllers.CreateStudent)
	r.PUT("/student/:id", controllers.UpdateStudentById)
	r.DELETE("/student/:id", controllers.DeleteStudentById)

	r.Run()
}
