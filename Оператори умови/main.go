package main

import "fmt"

func main() {
	v := 5

	switch v {
	case 1:
		fmt.Println("мій вибір BMW")

	case 2:
		fmt.Println("мій вибір Porsche")

	case 3,4:
		fmt.Println("мій вибір Mercedes")

	default:
		fmt.Println("мій вибір Audi")
	}
}
