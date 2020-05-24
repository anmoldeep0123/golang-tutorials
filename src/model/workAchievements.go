package model

import (
	"encoding/json"
	"fmt"
	"golangtuts/commonUtils"
	"golangtuts/constants"
	"golangtuts/interfacez"
)

type WorkAchievements struct {
	Name     string
	Place    string
	Year     int
	Employee interfacez.Employee
}

func (achievements WorkAchievements) RecordAnAchievement() {
	var fileHandler = interfacez.FileHandlerImpl(constants.EMP_BASE_DIR + "/" + achievements.Employee.GetAlias())
	createDir(fileHandler)
	fileHandler = interfacez.FileHandlerImpl(constants.EMP_BASE_DIR + "/" + achievements.Employee.GetAlias() + "/achievements" + commonUtils.GetRandomNumberAsStr() + ".txt")
	writeAchievementToFile(achievements.Employee, fileHandler)
}

func (achievements WorkAchievements) FetchAllAchievements() {
	var fileHandler = interfacez.FileHandlerImpl(constants.EMP_BASE_DIR + "/" + achievements.Employee.GetAlias())
	var achvFiles []string
	achvFiles = fileHandler.ListFiles()
	for i := 0; i < len(achvFiles); i++ {
		fileHandler = interfacez.FileHandlerImpl(achvFiles[i])
		fmt.Println(fileHandler.ReadFileInMem())
	}
}

func (achievements WorkAchievements) EmployeeDetails() {
	fmt.Println(achievements.Employee.GetAlias())
}

func writeAchievementToFile(employee interfacez.Employee, fileHandler interfacez.FileHandler) {
	jsonEmpl, err := json.Marshal(employee)
	if err != nil {
		panic(err)
	}
	fileHandler.WriteContent(string(jsonEmpl))
}

func createDir(fileHandler interfacez.FileHandler) {
	fileHandler.CreateDirectory()
}
