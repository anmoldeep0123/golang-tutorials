package main

import (
	"golangtuts/fileMgmt"
	"golangtuts/interfacez"
)

func main() {
	readOperationFromInterface();
}

func readOperationFromInterface() {
	filePath := "/Users/adeep/Documents/personal/jcg/May/import/info.txt"
	fileHandlerImpl := interfacez.FileHandlerImpl(filePath)
	var fileHandler interfacez.FileHandler
	fileHandler = fileHandlerImpl
	fileHandler.ReadFileInMem(filePath)
}

func readOperations() {
	fileMgmt.ReadFileInMem("/Users/adeep/Documents/personal/jcg/May/import/info.txt")
	fileMgmt.ReadInSmallChunks("/Users/adeep/Documents/personal/jcg/May/import/info.txt")
	fileMgmt.ReadFileLineByLine("/Users/adeep/Documents/personal/jcg/May/import/info.txt")
}

func writeOperations() {
	fileMgmt.CreateFileInResources("/Users/adeep/workspace/icode/goAdee/resources/sample.txt", "Hello World - Sample File\n")
}
