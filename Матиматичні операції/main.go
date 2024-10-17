package main

import "fmt"

func main() {
	fmt.Println(5-100, "віднімання")
	fmt.Println(5+15, "додавання")
	fmt.Println(2*100, "множення")

	k := 5
	v := 55

	if k == 0 {
		fmt.Println("на нуль ділити не можна")
	} else {
		fmt.Println(v/k, "ділення")
	}
}