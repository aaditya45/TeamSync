package service

import (
	"fmt"

	"gitlab.com/pragmaticreviews/golang-gin-poc/config"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	config.Hello()
	db = config.GetDBInstance()
	db.AutoMigrate(&entity.Employee{})
}

type EmployeeService interface {
	Add(entity.Employee) entity.Employee
	GetAll() []entity.Employee
}

type employeeService struct {
	employees []entity.Employee
}

func NewEmployee() EmployeeService {
	return &employeeService{}
}

func (service *employeeService) Add(employee entity.Employee) entity.Employee {
	fmt.Println("employee added in data")
	service.employees = append(service.employees, employee)
	//db.NewRecord(employee)
	db.Create(&employee)
	return employee
}

func (service *employeeService) GetAll() []entity.Employee {
	var data []entity.Employee
	fmt.Println("list of employees from data")
	res := db.Find(&data)
	fmt.Println(res.RowsAffected)
	return data
}
