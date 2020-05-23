package interfacez

import (
	"golangtuts/fileMgmt"
)

type FileHandlerImpl string

func (filePath FileHandlerImpl) ReadFileInMem() {
	fileMgmt.ReadFileInMem(string(filePath))
}

func (filePath FileHandlerImpl) ReadInSmallChunks() {
	fileMgmt.ReadInSmallChunks(string(filePath))
}

func (filePath FileHandlerImpl) ReadFileLineByLine() {
	fileMgmt.ReadFileLineByLine(string(filePath))
}

func (filePath FileHandlerImpl) CreateFileAndWriteContent(fileContent string) {
	fileMgmt.CreateFileInResources(string(filePath), fileContent)
}
