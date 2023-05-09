package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"gitlab.com/pragmaticreviews/golang-gin-poc/config"
	"gitlab.com/pragmaticreviews/golang-gin-poc/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/middlewares"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
	"gorm.io/gorm"
)

var (
	employeeService      service.EmployeeService         = service.NewEmployee()
	employeeController   controller.EmployeeController   = controller.NewEmployee(employeeService)
	departmentService    service.DepartmentService       = service.NewDepartment()
	departmentController controller.DepartmentController = controller.NewDepartment(departmentService)
	db                   *gorm.DB
)

func setupLogOutputInFile() {
	res, _ := os.Create("loggerFile.log")
	gin.DefaultWriter = io.MultiWriter(res, os.Stdout)
}

func init() {
	db = config.GetDBInstance()
}

func main() {
	setupLogOutputInFile()
	server := gin.New()

	//server.Static("/css", "./templates/css")
	//server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.SystemAuth(), gindump.Dump())

	server.GET("/getEmployees", func(ctx *gin.Context) {
		fmt.Println("getAll employee request")
		ctx.JSON(200, employeeController.GetAll())
	})

	server.POST("/addEmployee", func(ctx *gin.Context) {

		fmt.Println("employee added request")
		err := employeeController.Add(ctx)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(200, gin.H{"message": "data is added"})
		}

	})

	server.GET("/getDepartments", func(ctx *gin.Context) {
		fmt.Println("getAll department request")
		ctx.JSON(200, departmentController.GetAll())
	})

	server.POST("/addDepartment", func(ctx *gin.Context) {
		fmt.Println("department added request")
		err := departmentController.Add(ctx)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(200, gin.H{"message": "data is added"})
		}
	})

	server.Run(":8081")

}
