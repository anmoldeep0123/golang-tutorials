package fileMgmt

import (
	"fmt"
	"os"
)

func CreateDirectory(dirPath string) {
	if !dirExists(dirPath) {
		err := os.Mkdir(dirPath, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func CreateFileInResources(filePathName string, dataToWrite string) {
	byteArrayToWrite := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	multiLineStr := []string{"\nWelcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}
	filePtr, err := os.Create(filePathName)
	if err != nil {
		fmt.Println("CreateFileInResources ", err)
		return
	}
	writeOverWriteFile(filePtr, dataToWrite)
	writeBytesToFile(filePtr, byteArrayToWrite)
	writeLineByLine(filePtr, multiLineStr)
	err = filePtr.Close()
	if err != nil {
		fmt.Println("CreateFileInResources ", err)
		return
	}
	newLine := "File handling is easy."
	OpenAndAppendToFile(filePathName, newLine)
}

func OpenAndAppendToFile(fileName string, data string) {
	if !fileExists(fileName) {
		createNewFile(fileName)
	}
	filePtr, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("OpenAndAppendToFile ", err)
		return
	}
	_, err = fmt.Fprintln(filePtr, data)
	if err != nil {
		fmt.Println(err)
		filePtr.Close()
		return
	}
	err = filePtr.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("OpenAndAppendToFile: Data appended successfully")
}

func writeOverWriteFile(filePtr *os.File, data string) {
	bytezz, err := filePtr.WriteString(data)
	if err != nil {
		fmt.Println(err)
		filePtr.Close()
		return
	}
	fmt.Println(bytezz, "Bytes written successfully")
}

func writeBytesToFile(filePtr *os.File, byteArrayToWrite []byte) {
	bytezz, err := filePtr.Write(byteArrayToWrite)
	if err != nil {
		fmt.Println(err)
		filePtr.Close()
		return
	}
	fmt.Println(bytezz, "writeBytesToFile : bytes written successfully")
}

func writeLineByLine(filePtr *os.File, data []string) {
	for _, v := range data {
		fmt.Fprintln(filePtr, v)
	}
	fmt.Println("writeLineByLine : Data written successfully")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func dirExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func createNewFile(filePathName string) {
	_, err := os.Create(filePathName)
	if err != nil {
		fmt.Println("createNewFile ", err)
		return
	}
}
