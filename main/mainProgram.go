package main

import (
	L "../files"
)

func main() {
	L.ReadFileInMem("/Users/adeep/Documents/personal/jcg/May/import/info.txt")
	L.ReadInSmallChunks("/Users/adeep/Documents/personal/jcg/May/import/info.txt")
	L.ReadFileLineByLine("/Users/adeep/Documents/personal/jcg/May/import/info.txt")
}
