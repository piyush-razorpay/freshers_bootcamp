package main

import (
	"github.com/Gandhi24/api-practice/controllers"
	"github.com/Gandhi24/api-practice/repositories"
	"github.com/Gandhi24/api-practice/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	studentRepository repositories.StudentRepository = repositories.NewStudentRepository()
	studentService    services.StudentService        = services.New(studentRepository)
	studentController controllers.StudentController  = controllers.New(studentService)
)

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "I'm alive! Are you?"})
	})

	router.GET("/students", func(ctx *gin.Context) {
		ctx.JSON(200, studentController.FindAll())
	})

	router.POST("/students", func(ctx *gin.Context) {
		err := studentController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
		}

	})

	router.PUT("/students/:id", func(ctx *gin.Context) {
		err := studentController.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
		}

	})

	router.DELETE("/students/:id", func(ctx *gin.Context) {
		err := studentController.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
		}

	})

	router.Run()
}
