package utils

import (
	"golangtuts/interfacez"
	"golangtuts/model"
	"fmt"
)

func GetEmployee(employeeRole string) interfacez.Employee {
	var employee interfacez.Employee
	if employeeRole == "Manager" {
		fmt.Println("Constructing a Manager")
		employee = model.Manager{
			FirstName:   "Anmol",
			LastName:    "Deep",
			Role:        "Manager",
			Salary:      98415.70,
			LeavesTaken: 3,
		}
	} else if employeeRole == "HR" {
		fmt.Println("Constructing an HR")
		employee = model.HR{
			FirstName:   "Smrita",
			LastName:    "Irani",
			Role:        "Human Resources",
			Salary:      44500.15,
			LeavesTaken: 11,
		}
	} else {
		fmt.Println("Constructing an Operations")
		employee = model.Operations{
			FirstName:   "Ram",
			LastName:    "Reddy",
			Role:        "Operations",
			Salary:      25000.50,
			LeavesTaken: 6,
		}
	}
	return employee
}
