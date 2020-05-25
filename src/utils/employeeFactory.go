package utils

import (
	"fmt"
	"golangtuts/interfacez"
	"golangtuts/model"
)

func GetEmployee(employeeRole string) interfacez.Employee {
	var employee interfacez.Employee
	if employeeRole == "Manager" {
		fmt.Println("GetEmployee : Constructing a Manager")
		employee = model.Manager{
			Alias:       "adee",
			FirstName:   "Anmol",
			LastName:    "Deep",
			Role:        "Manager",
			Salary:      98415.70,
			LeavesTaken: 3,
		}
	} else if employeeRole == "HR" {
		fmt.Println("GetEmployee : Constructing an HR")
		employee = model.HR{
			Alias:       "sirani",
			FirstName:   "Smrita",
			LastName:    "Irani",
			Role:        "Human Resources",
			Salary:      44500.15,
			LeavesTaken: 11,
		}
	} else {
		fmt.Println("GetEmployee : Constructing an Operations")
		employee = model.Operations{
			Alias:       "rredyy",
			FirstName:   "Ram",
			LastName:    "Reddy",
			Role:        "Operations",
			Salary:      25000.50,
			LeavesTaken: 6,
		}
	}
	return employee
}
