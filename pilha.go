package main

import (
	"errors"
	"fmt"
)

func main() {
	pilha := Pilha{}
	fmt.Println("Pilha vazia?", pilha.EstaVazia())
	fmt.Printf("Tamanho da pilha: %d\n", pilha.GetTamanho())

	pilha.Empilhar("Go")
	pilha.Empilhar(3.14)
	pilha.Empilhar(5)

	fmt.Println("Pilha vazia?", pilha.EstaVazia())
	fmt.Printf("Tamanho da pilha: %d\n", pilha.GetTamanho())

	for !pilha.EstaVazia() {
		valor, _ := pilha.Desempilhar()
		fmt.Println("Valor =>", valor)
		fmt.Println("Pilha vazia?", pilha.EstaVazia())
		fmt.Printf("Tamanho da pilha: %d\n", pilha.GetTamanho())
	}

	_, err := pilha.Desempilhar()
	if err != nil {
		fmt.Println("erro:", err)
	}
}

type Pilha struct {
	valores []interface{}
}

func (pilha Pilha) GetTamanho() int {
	return len(pilha.valores)
}

func (pilha Pilha) EstaVazia() bool {
	return pilha.GetTamanho() == 0
}

func (pilha *Pilha) Empilhar(element interface{}) {
	pilha.valores = append(pilha.valores, element)
}

func (pilha *Pilha) Desempilhar() (interface{}, error) {
	if pilha.EstaVazia() {
		return nil, errors.New("Pilha vazia")
	}
	tamanhoPilha := pilha.GetTamanho()
	retorno := pilha.valores[tamanhoPilha-1]
	pilha.valores = pilha.valores[:tamanhoPilha-1]
	return retorno, nil
}
