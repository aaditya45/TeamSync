package controller

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

type DepartmentController interface {
	GetAll() []entity.Department
	Add(ctx *gin.Context) error
}

type departmentController struct {
	service service.DepartmentService
}

func NewDepartment(service service.DepartmentService) DepartmentController {
	return &departmentController{
		service: service,
	}
}

func (c *departmentController) GetAll() []entity.Department {
	fmt.Println("getAll department at controller")
	return c.service.GetAll()
}

func (c *departmentController) Add(ctx *gin.Context) error {
	fmt.Println("department added at controller")
	var department entity.Department
	err := ctx.ShouldBindJSON(&department)
	if err != nil {
		return err
	}
	department.Title = strings.Trim(department.Title, " ")
	if len(department.Title) == 0 {
		return errors.New("size of title is 0")
	}
	if department.DepartmentCode <= 0 {
		return errors.New("department code should be greater than 0")
	}
	c.service.Add(department)
	return nil
}
