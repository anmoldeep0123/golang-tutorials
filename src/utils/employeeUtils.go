package utils

import (
	"fmt"
	"golangtuts/interfacez"
)

func CreateEmployees() []interfacez.Employee {
	const count = 5
	var employees [5]interfacez.Employee
	employeeType := [count]string{"HR", "Operations", "HR", "Manager", "HR"}
	for i := 0; i < count; i++ {
		employees[i] = createEmployee(employeeType[i])
		fmt.Println(employees[i])
	}
	return employees[:]
}

func CreateAndDescribeEmployee(employeeType string) {
	var employee interfacez.Employee = createEmployee(employeeType)
	employeeDetails(employee)
}

func employeeDetails(employee interfacez.Employee) {
	employee.DescribeEmployee()
}

func createEmployee(employeeType string) interfacez.Employee {
	return GetEmployee(employeeType)
}
