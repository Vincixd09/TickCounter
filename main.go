package main

import (
	"fmt"
)

func menu() {
	fmt.Print("1: Change of Seconds by ticks\n")
	fmt.Print("2: Chanve of Minute by ticks\n")
	fmt.Print("3: Change of Horas by ticks\n")
	fmt.Print("0: Exit\n")
}

func calc(fn func(float64) float64) {
	fmt.Print("\033[3;J\033[H\033[2J")
	fmt.Println("Insert number")
	var num float64
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := fn(num)
	fmt.Print("\033[3;J\033[H\033[2J")
	fmt.Print(result, "\n")
	main()
}

func seg(a float64) float64 {
	return a * 20
}

func min(a float64) float64 {
	return a * 1200
}

func hora(a float64) float64 {
	return a * 72000
}

func main() {
	menu()
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println(err)
		return
	}

	if num == 1 {
		calc(seg)
	} else if num == 2 {
		calc(min)
	} else if num == 3 {
		calc(hora)
	} else if num == 0 {
		fmt.Println("Exit...")
		return
	}
}
