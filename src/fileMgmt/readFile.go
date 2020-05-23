package fileMgmt

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadFileInMem(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error while reading file ", err)
		return
	}
	fmt.Println("Contents of the file are : ", string(data))
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
			fmt.Println("Error reading file ", err)
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