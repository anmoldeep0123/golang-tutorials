package main

import (
	L "../files"
)

func main() {
	writeOperations()
}

func readOperations() {
	L.ReadFileInMem("/Users/adeep/Documents/personal/jcg/May/import/info.txt")
	L.ReadInSmallChunks("/Users/adeep/Documents/personal/jcg/May/import/info.txt")
	L.ReadFileLineByLine("/Users/adeep/Documents/personal/jcg/May/import/info.txt")
}

func writeOperations() {
	L.CreateFileInResources("/Users/adeep/workspace/icode/goAdee/resources/sample.txt" , "Hello World - Sample File\n");
}
