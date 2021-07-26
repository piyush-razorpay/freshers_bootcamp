package services

import (
	"github.com/Gandhi24/api-practice/models"
	"github.com/Gandhi24/api-practice/repositories"
)

type StudentService interface {
	Save(models.Student) models.Student
	Update(models.Student) models.Student
	Delete(models.Student) models.Student
	FindAll() []models.Student
}

type studentService struct {
	repository repositories.StudentRepository
}

func New(repo repositories.StudentRepository) StudentService {
	return &studentService{
		repository: repo,
	}
}

func (s *studentService) Save(student models.Student) models.Student {
	s.repository.Save(student)
	return student
}

func (s *studentService) Update(student models.Student) models.Student {
	s.repository.Update(student)
	return student
}

func (s *studentService) Delete(student models.Student) models.Student {
	s.repository.Delete(student)
	return student
}

func (s *studentService) FindAll() []models.Student {
	students := s.repository.FindAll()
	return students
}
