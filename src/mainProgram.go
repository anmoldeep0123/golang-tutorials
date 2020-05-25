package main

import (
	"fmt"
	"golangtuts/driver"
	"golangtuts/fileMgmt"
	"golangtuts/interfacez"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	go driver.CacheExamples(ch1)
	fmt.Println("Channel 1 Output", <-ch1)
	go driver.GoCompositions(ch2)
	fmt.Println("Channel 2 Output", <-ch2)
	go driver.DealWithEmployees(ch3)
	fmt.Println("Channel 3 Output", <-ch3)

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
