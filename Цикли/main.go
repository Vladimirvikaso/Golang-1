package main

import "fmt"

func main() {
	Names := []string{"Підтягування", "Присідання", "Віджимання"}
	for _, Name := range Names {
		fmt.Println("Прийшов час вправи:", Name)
		for count := 0; count < 10; count++ {
			fmt.Println(Name, count+1)
		}
	}
}