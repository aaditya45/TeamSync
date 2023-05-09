package entity

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Title          string `json : "title" binding : "min=2,max=10"`
	DepartmentCode int    `json : "departmentCode" binding : "gte=1"`
}

type Employee struct {
	gorm.Model
	EmployeeName   string     `gorm:"" json : "employeeName" binding : "required"`
	EmployeeId     int        `json : "employeeId" binding : "required"`
	Designation    string     `json : "designation"`
	Age            int8       `json : "age" binding : "gte=1,lte=130"`
	Email          string     `json : "email" binding : "required,email"`
	DepartmentInfo Department `gorm:"embedded" json : "departmentInfo" binding : "required"`
}
