package interfacez

import (
	"golangtuts/fileMgmt"
)

type FileHandlerImpl string

func (filePath FileHandlerImpl) ReadFileInMem() string {
	return fileMgmt.ReadFileInMem(string(filePath))
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

func (filePath FileHandlerImpl) CreateDirectory() {
	fileMgmt.CreateDirectory(string(filePath))
}

func (filePath FileHandlerImpl) WriteContent(fileContent string) {
	fileMgmt.OpenAndAppendToFile(string(filePath), fileContent)
}

func (filePath FileHandlerImpl) ListFiles() []string {
	return fileMgmt.ListFiles(string(filePath))
}
