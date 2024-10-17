package main

import "fmt"

func main() {
	var Vova, Danil, Vlad int

	Vova = 3
	Danil = 2
	Vlad = 10

	hinkali := Vova + Danil + Vlad
	order := fmt.Sprintf("Замовлення - %s, кількість - %d", "хінкалі", hinkali)

	fmt.Println(order)
}