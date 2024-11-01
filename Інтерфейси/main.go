package main

import "fmt"

type boiler interface {
	boil()
}

type kettle struct {}

func (kettle) boil() {
	fmt.Println("чайник: вода закипячена")
}

type pot struct {}

func(pot) boil() {
	fmt.Println("баняк: вода закипячена")
}

func makeMivina(b boiler) {
	fmt.Println("взяти тарілку")
	fmt.Println("покласти мівіну в тарілку")

	b.boil()

	fmt.Println("залити мівіну водою")
	fmt.Println("накрити кришкою", "\n")
}

func main() {
	makeMivina(pot{})
}