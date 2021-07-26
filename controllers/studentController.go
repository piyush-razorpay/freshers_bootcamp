package controllers

import (
	"github.com/Gandhi24/api-practice/models"
	"github.com/Gandhi24/api-practice/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

var validate *validator.Validate

type StudentController interface {
	FindAll() []models.Student
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type Controller struct {
	service services.StudentService
}

func New(studentService services.StudentService) StudentController {
	validate = validator.New()
	return &Controller{
		service: studentService,
	}
}

func (c *Controller) Save(ctx *gin.Context) error {
	var student models.Student
	err := ctx.ShouldBindJSON(&student)
	if err != nil {
		return err
	}
	err = validate.Struct(student)
	if err != nil {
		return err
	}
	c.service.Save(student)
	return nil
}

func (c *Controller) Update(ctx *gin.Context) error {
	var student models.Student
	err := ctx.ShouldBindJSON(&student)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	student.ID = id

	err = validate.Struct(student)
	if err != nil {
		return err
	}
	c.service.Update(student)
	return nil
}

func (c *Controller) Delete(ctx *gin.Context) error {
	var student models.Student
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	student.ID = id
	c.service.Delete(student)
	return nil
}

func (c *Controller) FindAll() []models.Student {
	return c.service.FindAll()
}
