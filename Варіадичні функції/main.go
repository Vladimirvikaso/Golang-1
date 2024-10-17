package main

import "fmt"

var menu = []string{"хінкалі", "хачапурі", "харчо", "чібупелі", "шаурма"}

func main() {
	var Liubomyr, Yura, Bohdan, Oleh uint
	Liubomyr = 2
	Yura = 4
	Bohdan = 4
	Oleh = 15

	result, err := order("шаурма", Liubomyr, Yura, Bohdan, Oleh)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func order(dish string, amount ...uint) (string, error) {
	var finalAmount uint
	for _, item := range amount {
		finalAmount += item
	}

	for _, item := range menu {
		if dish == item {
			return fmt.Sprintf("Замовлення = %s, кількість = %d", dish, finalAmount), nil
		}
	}

	return "", fmt.Errorf("Блюдо не знайдено")
}