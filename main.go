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

func seg() {
	fmt.Print("How much do you want to change from seconds to ticks?\n")

	var seg int
	_, err := fmt.Scanln(&seg)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := seg * 20
	fmt.Print(result, "\n")
	main()
}

func min() {
	fmt.Print("How much do you want to change from seconds to ticks: \n")
	
	var min int
	_, err := fmt.Scanln(&min)
	if err != nil {
		fmt.Println(err)
		return
	}
			
	result := min * 1200
	fmt.Print(result, "\n")
	main()
}

func hora() {
	fmt.Print("How much do you want to change from hours to ticks: \n")

	var hora int
	_, err := fmt.Scanln(&hora)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := hora * 72000
	fmt.Print(result, "\n")
	main()
}

func main() {
	menu()
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println("Error al leer el numero: ", err)
		return
	}


	if num == 1 {
		seg()
	}else if num == 2 {
		min()
	}else if num == 3 {
		hora()
	}else if num == 0 {
		fmt.Println("Saliendo...")
		return
	}
}
