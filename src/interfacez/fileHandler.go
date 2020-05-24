package interfacez

type FileHandler interface {
	ReadFileInMem() string
	ReadInSmallChunks()
	ReadFileLineByLine()
	CreateFileAndWriteContent(fileContent string)
	CreateDirectory()
	WriteContent(fileContent string)
	ListFiles() []string
}
