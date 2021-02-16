package controllers

import (
	"gowebservice/config"
	"gowebservice/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Student

	if err := config.DB.Find(&students).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, students)
}

func GetStudentById(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student
	if err := config.DB.Where("id = ?", id).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var student models.Student

	fullname := c.PostForm("fullname")
	age := c.PostForm("age")
	ageNumber, _ := strconv.Atoi(age)
	batch := c.PostForm("batch")

	student.Fullname = fullname
	student.Age = ageNumber
	student.Batch = batch

	if err := config.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't create!"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func UpdateStudentById(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student
	var newStudent models.Student

	if err := config.DB.Where("id = ?", id).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	fullname := c.PostForm("fullname")
	age := c.PostForm("age")
	ageNumber, _ := strconv.Atoi(age)
	batch := c.PostForm("batch")

	newStudent.Fullname = fullname
	newStudent.Age = ageNumber
	newStudent.Batch = batch

	if err := config.DB.Model(&student).Updates(newStudent).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't update!"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func DeleteStudentById(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student

	if err := config.DB.Where("id = ?", id).Delete(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
}
