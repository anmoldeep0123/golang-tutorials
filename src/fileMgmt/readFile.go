package fileMgmt

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func ListFiles(filePath string) []string {
	var files []string
	filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func ReadFileInMem(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFileInMem: Error while reading file ", err)
		return "nil"
	}
	return string(data)
}

func ReadInSmallChunks(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 255)
	for {
		n, err := r.Read(b)
		if err != nil {
			fmt.Println("ReadInSmallChunks: Error reading file ", err)
			break
		}
		fmt.Println(string(b[0:n]))
	}
}

func ReadFileLineByLine(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
