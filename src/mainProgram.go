package main

import (
	"fmt"
	"golangtuts/fileMgmt"
	"golangtuts/interfacez"
	"golangtuts/model"
)

func main() {
	structImplementInterface()
}

func structImplementInterface() {
	var employee interfacez.Employee = model.Manager{
		FirstName:   "Sam",
		LastName:    "Adolf",
		Salary:      98415.70,
		LeavesTaken: 3,
	}
	fmt.Println(employee.SalaryForLeavesTaken())
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
