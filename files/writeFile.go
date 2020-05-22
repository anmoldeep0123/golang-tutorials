package files

import (
	"os"
	"fmt"
)

func CreateFileInResources(filePathName string , dataToWrite string) {  
	byteArrayToWrite := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	multiLineStr := []string{"\nWelcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}
	filePtr, err := os.Create(filePathName)
    if err != nil {
        fmt.Println(err)
        return
	}
	writeOverWriteFile(filePtr , dataToWrite)
	writeBytesToFile(filePtr,byteArrayToWrite)
	writeLineByLine(filePtr , multiLineStr)
	err = filePtr.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
    newLine := "File handling is easy."
    openAndAppendToFile(filePathName , newLine)
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
    fmt.Println(bytezz, "bytes written successfully")
}

func writeLineByLine(filePtr *os.File, data []string) {
    for _, v := range data {
        fmt.Fprintln(filePtr, v)
    }
    fmt.Println("Data written successfully")
}

func openAndAppendToFile(fileName string, data string) {
    filePtr, err := os.OpenFile(fileName , os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
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
    fmt.Println("Data appended successfully")
}
