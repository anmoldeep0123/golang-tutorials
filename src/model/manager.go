package model

import (
	"fmt"
)

type Manager struct {
	Alias       string
	FirstName   string
	LastName    string
	Role        string
	Salary      float32
	LeavesTaken int
}

func (e Manager) SalaryForLeavesTaken() float32 {
	return e.Salary * float32(e.LeavesTaken) / 31
}

func (e Manager) DescribeEmployee() {
	fmt.Println("Employee ", e.Role, " Details ", e.FirstName, " ", e.LastName, " , ", "Draws ", e.Salary, " Leaves Taken ", e.LeavesTaken)
}

func (e Manager) GetAlias() string {
	return e.Alias
}
