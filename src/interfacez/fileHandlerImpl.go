package interfacez

import (
	"golangtuts/fileMgmt"
	"fmt"
)

type FileHandlerImpl string

func (filePath FileHandlerImpl) ReadFileInMem(path string) {
	fmt.Println("Path to read File is : " , path)
	fileMgmt.ReadFileInMem(string(filePath))
}

func (filePath FileHandlerImpl) ReadInSmallChunks(path string) {
	fileMgmt.ReadInSmallChunks(string(filePath))
}

func (filePath FileHandlerImpl) ReadFileLineByLine(path string) {
	fileMgmt.ReadFileLineByLine(string(filePath))
}