package repositories

import (
	"fmt"
	"github.com/Gandhi24/api-practice/models"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

type StudentRepository interface {
	Save(Student models.Student)
	Update(Student models.Student)
	Delete(Student models.Student)
	FindAll() []models.Student
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func NewStudentRepository() StudentRepository {
	db, err := gorm.Open("mysql", "root:abraham.lincoln##1@tcp(localhost:3306)/practice_api?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&models.Student{})
	return &database{
		connection: db,
	}
}

func (db *database) Save(Student models.Student) {
	db.connection.Create(&Student)
}

func (db *database) Update(Student models.Student) {
	db.connection.Save(&Student)
}

func (db *database) Delete(Student models.Student) {
	db.connection.Delete(&Student)
}

func (db *database) FindAll() []models.Student {
	var Students []models.Student
	db.connection.Find(&Students)
	return Students
}
