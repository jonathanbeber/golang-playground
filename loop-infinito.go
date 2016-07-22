package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	iteracoes := 0

	for {
		iteracoes++

		randomNumber := rand.Intn(4200)
		fmt.Println(randomNumber)
		if randomNumber%42 == 0 {
			break
		}
	}
	fmt.Printf("saindo após %d iterações!", iteracoes)
}
