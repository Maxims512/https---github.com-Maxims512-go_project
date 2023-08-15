package main

import (
	"fmt"
	"os"
)

func openFile() {
	file, _ := os.Open("test.txt")
	data, _ := os.ReadFile(file.Name())
	fmt.Println(string(data))
}

func main() {
	openFile()
}
