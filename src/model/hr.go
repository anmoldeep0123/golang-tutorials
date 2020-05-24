package model

import (
	"fmt"
)

type HR struct {
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
	fmt.Println("Employee ", e.Role, " Details ", e.FirstName, " ", e.LastName, " , ", "Draws ", e.Salary, " Leaves Taken ", e.LeavesTaken)
}
