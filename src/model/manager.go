package model

import ()

type Manager struct {
	FirstName   string
	LastName    string
	Salary      float32
	LeavesTaken int
}

func (e Manager) SalaryForLeavesTaken() float32 {
	return e.Salary * float32(e.LeavesTaken) / 31
}
