package main

import (
	"fmt"
	"golangtuts/fileMgmt"
	"golangtuts/interfacez"
	"golangtuts/model"
	"golangtuts/utils"
)

func main() {
	recordAchievements()
}

func recordAchievements() {
	var emp interfacez.Employee = utils.CreateEmployee("HR")
	var achv interfacez.Achievements = model.WorkAchievements{
		Name:     "Employee Of The Month",
		Place:    "Bangalore",
		Year:     2019,
		Employee: emp,
	}
	achv.RecordAnAchievement()
	achv.FetchAllAchievements()
	achv.EmployeeDetails()
}

func dealWithEmployees() {
	var employees []interfacez.Employee = utils.CreateEmployees()
	utils.CreateAndDescribeEmployee("HR")
	fmt.Println(employees)
}

func readWriteOperationUsingInterface() {
	readFilePath := "/Users/adeep/Documents/personal/jcg/May/import/info.txt"
	writeFilePath := "/Users/adeep/workspace/icode/goAdee/resources/sample.txt"
	writeContent := "Hello World - Sample File\n"

	readFileHandlerImpl := interfacez.FileHandlerImpl(readFilePath)
	writeFileHandlerImpl := interfacez.FileHandlerImpl(writeFilePath)
	var fileHandler interfacez.FileHandler

	fileHandler = readFileHandlerImpl
	fileHandler.ReadFileInMem()
	fileHandler.ReadInSmallChunks()
	fileHandler.ReadFileLineByLine()

	fileHandler = writeFileHandlerImpl
	fileHandler.CreateFileAndWriteContent(writeContent)
}

func readOperations() {
	fileMgmt.ReadFileInMem("/Users/adeep/Documents/personal/jcg/May/import/info.txt")
	fileMgmt.ReadInSmallChunks("/Users/adeep/Documents/personal/jcg/May/import/info.txt")
	fileMgmt.ReadFileLineByLine("/Users/adeep/Documents/personal/jcg/May/import/info.txt")
}

func writeOperations() {
	fileMgmt.CreateFileInResources("/Users/adeep/workspace/icode/goAdee/resources/sample.txt", "Hello World - Sample File\n")
}
