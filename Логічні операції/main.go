package main

import "fmt"

func main() {
	shopOpen := true
	chef := true
	bread := true
	iischef := false

	if (shopOpen && chef && bread) || iischef {
		fmt.Println("Я можу купити хліб")
	}
}