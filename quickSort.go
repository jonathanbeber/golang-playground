package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Uso: %s <numeros>\n", os.Args[0])
		os.Exit(1)
	}
	valoresOrigem := os.Args[1:]
	resultado := make([]int, len(valoresOrigem))

	for indice, value := range valoresOrigem {
		valorConvertido, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("Falha ao converter \"%s\"\n", value)
			os.Exit(1)
		}
		resultado[indice] = valorConvertido
	}

	fmt.Println(quickSort(resultado))
}

func quickSort(numeros []int) []int {
	if len(numeros) <= 1 {
		return numeros
	}

	resultado := make([]int, len(numeros))
	copy(resultado, numeros)

	indicePivo := len(resultado) / 2
	pivo := resultado[indicePivo]
	resultado = append(resultado[:indicePivo], resultado[indicePivo+1:]...)

	menores, maiores := particionar(resultado, pivo)

	return append(
		append(quickSort(menores), pivo),
		quickSort(maiores)...)
}

func particionar(numeros []int, pivo int) (menores []int, maiores []int) {
	for _, numero := range numeros {
		if numero <= pivo {
			menores = append(menores, numero)
		} else {
			maiores = append(maiores, numero)
		}
	}

	return menores, maiores
}
