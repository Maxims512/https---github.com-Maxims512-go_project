package main

import (
	"fmt"
	"os"
	//"strconv"
)

func openFile() string {
	file, _ := os.Open("voprosi.txt")
	data, _ := os.ReadFile(file.Name())
	return string(data)
}

func give_vopr_otv(text string) (array[100]rune) {
	var schet int
	var perem int
	//var otvet rune
	for _, i := range text {
		if i == 13 || i == 10{
			schet++	
		}
		if schet == 8{
			if i != 10{
				array[perem] = i
				perem++
				schet=-4
			}

		}
	}
	return array
}

//func give_vopr(text string) (array[100]string)


func main() {
	data := openFile()
	vopr_otvet := give_vopr_otv(data)
	for i := 0; i < 100; i++{
        fmt.Println(i+1, string(vopr_otvet[i]))
    }
}
