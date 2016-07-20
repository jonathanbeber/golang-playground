package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Uso: %s <palavras>\n", os.Args[0])
		os.Exit(1)
	}
	imprimir(colherEstatisticas(os.Args[1:]))
}

func colherEstatisticas(palavras []string) map[string]int {
	retorno := make(map[string]int)
	for _, palavra := range palavras {
		letra := strings.ToUpper(string(palavra[0]))
		_, achou := retorno[letra]
		if achou {
			retorno[letra]++
		} else {
			retorno[letra] = 1
		}
	}
	return retorno
}

func imprimir(mapa map[string]int) {
	for letra, contador := range mapa {
		fmt.Printf("%s ->\t%d\n", letra, contador)
	}
}
