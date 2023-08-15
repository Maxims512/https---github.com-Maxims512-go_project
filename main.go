package main

import (
	"fmt"
	"os"
)

func openFile() string {
	file, _ := os.Open("test.txt")
	data, _ := os.ReadFile(file.Name())
	return string(data)
}

func give_vopr_otv(text string) (vopros_otvet []int) {
	i := 0
	perem := ""
	fmt.Println(len(text))
	//
}

func main() {
	data := openFile()
	voprOtv := give_vopr_otv(data)

}
