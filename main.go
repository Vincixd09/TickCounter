package main

import (
	"fmt"
)

func menu() {
	fmt.Print("1: Change of Seconds by ticks\n")
	fmt.Print("2: Chanve of Minute by ticks\n")
	fmt.Print("3: Change of Hours by ticks\n")
	fmt.Print("4: Change of Ticks by Seconds\n")
	fmt.Print("5: Change of Ticks by Minute\n")
	fmt.Print("6: Change of Ticks by Hours\n")
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

func main() {
	menu()
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch num {
	case 1:
		calc(func(f float64) float64 { return f * 20 })
	case 2:
		calc(func(f float64) float64 { return f * 1200 })
	case 3:
		calc(func(f float64) float64 { return f * 72000 })
	case 4:
		calc(func(f float64) float64 {return f / 20})
	case 5:
		calc(func(f float64) float64 {return f / 1200})
	case 6:
		calc(func(f float64) float64 {return f / 72000})
	case 0:
		fmt.Print("\033[3;J\033[H\033[2J")
		fmt.Println("Exit...")
		return
	}
}
