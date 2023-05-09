package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

type EmployeeController interface {
	GetAll() []entity.Employee
	Add(ctx *gin.Context) error
}

type employeeController struct {
	service service.EmployeeService
}

func NewEmployee(service service.EmployeeService) EmployeeController {
	return &employeeController{
		service: service,
	}
}

func (c *employeeController) GetAll() []entity.Employee {
	fmt.Println("getAll employee at controller")
	return c.service.GetAll()
}

func (c *employeeController) Add(ctx *gin.Context) error {
	fmt.Println("employee added at controller")
	var employee entity.Employee
	err := ctx.ShouldBindJSON(&employee)
	if err != nil {
		return err
	}
	c.service.Add(employee)
	return nil
}
