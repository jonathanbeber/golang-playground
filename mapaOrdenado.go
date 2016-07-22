package main

import "fmt"

func main() {
	quadrado := make(map[int]int, 15)
	var numbers []int
	for i := 1; i <= 15; i++ {
		quadrado[i] = i * i
		numbers = append(numbers, i)
	}
	for _, number := range numbers {
		fmt.Println(number, "^ 2 =", quadrado[number])
	}

}
