package model

import (
	"fmt"
)

type Operations struct {
	Alias       string
	FirstName   string
	LastName    string
	Role        string
	Salary      float32
	LeavesTaken int
}

func (e Operations) SalaryForLeavesTaken() float32 {
	return e.Salary * float32(e.LeavesTaken) / 31
}

func (e Operations) DescribeEmployee() {
	fmt.Println("Employee ", e.Role, " Details ", e.FirstName, " ", e.LastName, " , ", "Draws ", e.Salary, " Leaves Taken ", e.LeavesTaken)
}

func (e Operations) GetAlias() string {
	return e.Alias
}
