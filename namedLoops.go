package main

import "fmt"

func main() {

forzao:
	for index := 0; index < 10; index++ {
		fmt.Println("for i => ", index)

		switch index {
		case 2, 3:
			fmt.Println("é dois ou três")
			break
		case 5:
			fmt.Println("Quebrando tudo!")
			break forzao
		}

	}
}
