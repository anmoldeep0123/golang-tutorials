package interfacez

import ()

type FileHandler interface {
	ReadFileInMem()
	ReadInSmallChunks()
	ReadFileLineByLine()
	CreateFileAndWriteContent(fileContent string)   
}
