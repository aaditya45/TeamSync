package service

import (
	"fmt"

	//"github.com/jinzhu/gorm"
	"gitlab.com/pragmaticreviews/golang-gin-poc/config"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
)

func init() {
	db = config.GetDBInstance()
	db.AutoMigrate(&entity.Department{})

}

type DepartmentService interface {
	Add(entity.Department) entity.Department
	GetAll() []entity.Department
}

type departmentService struct {
	department []entity.Department
}

func NewDepartment() DepartmentService {
	return &departmentService{}
}

func (service *departmentService) Add(department entity.Department) entity.Department {
	fmt.Println("employee added in data")
	service.department = append(service.department, department)
	//db.NewRecord(employee)
	db.Create(&department)
	return department
}

func (service *departmentService) GetAll() []entity.Department {
	var data []entity.Department
	fmt.Println("list of employees from data")
	res := db.Find(&data)
	fmt.Println(res.RowsAffected)
	return data
}
