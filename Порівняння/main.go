package main

import "fmt"

func main() {
	cookiesPrice := 15
	myMoney := 15

	canByCookies := myMoney >= cookiesPrice
	fmt.Println("Чи можу я купити печиво:", canByCookies)

	cookiesRoshen := "Tasty"
	myCookies := "Tastyone"

	cookiesForMe := cookiesRoshen == myCookies
	fmt.Println("Те що треба?", cookiesForMe)

	cookiesNotMe := cookiesRoshen != myCookies
	fmt.Println("Не підходить?", cookiesNotMe)
}