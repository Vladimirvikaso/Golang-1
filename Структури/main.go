package main

import "fmt"

type wallet struct {
	creditCards []creditCard
}

type creditCard struct{
	number int
	expirationDate string
	cvvCode int
	holder string
}

func main() {
	VictoriaCreditCard := creditCard{
		number: 1111222233334444,
		expirationDate: "24 лютого 2022р",
		cvvCode: 230,
		holder: "Вікторія",
	}

	viktoriaMonoCreditCardNumber := creditCard{
		number: 1212131314141515,
		expirationDate: "01 січня 2020р",
		cvvCode: 222,
		holder: "Вікторія",
	}

	viktoriaPrivateCreditCardNumber := creditCard{
		number: 3333555577779999,
		expirationDate: "05 травня 2015р",
		cvvCode: 999,
		holder: "Вікторія",
	}

	myWallet := wallet{
		creditCards: []creditCard{
			VictoriaCreditCard,
			viktoriaMonoCreditCardNumber,
			viktoriaPrivateCreditCardNumber,
		},
	}

	payWithWallet(myWallet, 100)
	

	cardNumber := 1111222233334444
	cardExpirationDate := "24 лютого 2022р"
	cardCvvCode := 230
	cardHolderName := "Вікторія"
	amountToPay := 100

	pay(cardNumber, cardExpirationDate, cardCvvCode, cardHolderName, amountToPay,)

	cardNumberMono := 1212131314141515
	viktoriaMonoCardExpirationDate := "01 січня 2020р"
	viktoriaMonoCardCvvCode := 222
	viktoriaMonoCardHolderName := "Вікторія"

	pay(cardNumberMono, viktoriaMonoCardExpirationDate, viktoriaMonoCardCvvCode, viktoriaMonoCardHolderName, amountToPay,)

	cardNumberPrivat := 1490136812340990
	viktoriaPrivateCardExpirationDate := "05 травня 2015р"
	viktoriaPrivateCardCvvCode := 999
	viktoriaPrivateCardHolderName := "Вікторія"

	pay(cardNumberPrivat, viktoriaPrivateCardExpirationDate, viktoriaPrivateCardCvvCode, viktoriaPrivateCardHolderName, amountToPay,)

}



func payWithCreditCard(c creditCard, amount int) {
	fmt.Println("Інформація по картці")
	fmt.Println("Номер картки: ", c.number)
	fmt.Println("Дата створення: ", c.expirationDate)
	fmt.Println("Сvv: ", c.cvvCode)
	fmt.Println("І'мя власника(ця): ", c.holder)
	fmt.Println("Сумма переводу: ", amount)
}

func pay(cardNumber int, cardExpirationDate string, cardCvvCode int, cardHolderName string, amountToPay int) {

	fmt.Println("Інформація по картці")
	fmt.Println("Номер картки: ", cardNumber)
	fmt.Println("Дата створення: ", cardExpirationDate)
	fmt.Println("Сvv: ", cardCvvCode)
	fmt.Println("І'мя власника(ця): ", cardHolderName)
	fmt.Println("Сумма переводу: ", amountToPay)
}

func payWithWallet(w wallet, amount int) {
	for _, card := range w.creditCards {
		payWithCreditCard(card, amount)
	}
}