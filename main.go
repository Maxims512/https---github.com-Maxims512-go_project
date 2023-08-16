package main

import (
	"fmt"
	"os"
	"math/rand"
)

func openFile() string {
	file, _ := os.Open("voprosi.txt")
	data, _ := os.ReadFile(file.Name())
	return string(data)
}

func give_vopr_otv(text string) (array[100]rune) {
	var schet int
	var perem int
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

func give_vopr(text string) (array[100]string){
	var schet int
	var vopr string
	var perem int
	for _, i := range text {
		if schet<8 {
			vopr += string(i)
		} else{
			if perem==0{
				array[perem]=vopr
			} else {
				array[perem]=vopr[4:]
			}
			vopr=""
			perem++
			schet=-4
		}
		if i == 13 || i == 10{
			schet++	
		}
	}
	return array
}




func main() {
	data := openFile()
	vopr := give_vopr(data)
	otveti := give_vopr_otv(data)
	perem := ""
	fmt.Println("Вы открыли тест по безопасному владению оружием")
	fmt.Println("Вы можете сделать одну ошибку чтобы удачно сдать тест")
	fmt.Println("Для ответа на вопрос нажмите цифру правильного ответа на клавиатуре")
	fmt.Println("Нажмите Enter чтобы начать тест")
	fmt.Fscan(os.Stdin, perem) 
	fmt.Printf("\x1bc")
	polz_otvet := ""
	for {
		oshibk := 0
		for i := 0; i < 10; i++{
			random := rand.Intn(92)
			fmt.Println(vopr[random])
			otv := string(otveti[random])
			fmt.Fscan(os.Stdin, &perem) 
			if polz_otvet != otv{
				oshibk++
			}
			fmt.Printf("\x1bc")
		}
		if oshibk>1{
			fmt.Println("Вы не сдали тест")
		} else {
			fmt.Println("Вы сдали тест")
		}
		break
	}
	fmt.Fscan(os.Stdin, perem) 
}
