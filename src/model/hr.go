package model

import (
	"fmt"
)

type HR struct {
	Alias       string
	FirstName   string
	LastName    string
	Role        string
	Salary      float32
	LeavesTaken int
}

func (e HR) SalaryForLeavesTaken() float32 {
	return e.Salary * float32(e.LeavesTaken) / 31
}

func (e HR) DescribeEmployee() {
	fmt.Println("HR -> DescribeEmployee : Employee ", e.Role, " Details ", e.Alias, " - ", e.FirstName, " ", e.LastName, " , ", "Draws ", e.Salary, " Leaves Taken ", e.LeavesTaken)
}

func (e HR) GetAlias() string {
	return e.Alias
}
