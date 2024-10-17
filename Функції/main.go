package main

import "fmt"

var menu = []string{"хінкалі", "харчо", "хачапурі", "шаурма"}

func main() {
	result, err := order("шаурма", 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	result, err = order("ананасовий сок", 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func order(dish string, amount uint) (string, error){
	for _, item := range menu {
		if dish == item {
			return fmt.Sprintf("Замовлення = %s, кількість = %d", dish, amount), nil
		}
	}

	return "", fmt.Errorf("Блюдо не знайдено")
}
