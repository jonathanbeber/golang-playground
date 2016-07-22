package main

import "fmt"

type Estado struct {
	nome      string
	capital   string
	populacao int
}

func main() {
	estados := make(map[string]Estado, 6)
	estados["GO"] = Estado{"Goias", "Goiania", 2893923}
	fmt.Println(estados["GO"].nome)
}
