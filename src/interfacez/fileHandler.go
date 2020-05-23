package interfacez

import ()

type FileHandler interface {
	ReadFileInMem(path string)
	ReadInSmallChunks(path string)
	ReadFileLineByLine(path string)
}
