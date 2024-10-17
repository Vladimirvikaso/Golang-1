package main

import "fmt"

func main() {
	age := 0
	fmt.Println("Початковий вік автора", age)

	age = 25
	fmt.Println("Вік автора", age)

	age = 26
	fmt.Println("Змінений вік", age)

	date := "27 травня 1998р"
	fmt.Println(date, "один")
	fmt.Println(date, "два")
	fmt.Println(date, "три")

	date = "27 травня 2024"
	fmt.Println(date, "чотири")
	fmt.Println(date, "сім")
}

