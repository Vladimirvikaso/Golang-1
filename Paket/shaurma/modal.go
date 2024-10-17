package shaurma

import (
	"fmt"
	"recipes/cabbage"
	"recipes/meat"
)

var Name = "Шаурма"

type Shaurma struct {
	Name string
	FreshCabbage cabbage.Cabbage
	FriedMeat meat.Meat

}

func EatUp(s Shaurma) {
	fmt.Println(fmt.Printf("Смачного! %s зайшла як додому. Треба взяти ще одну!", s.Name))
	fmt.Println(fmt.Printf("Складники: %s вагою %d", s.FriedMeat.Name, s.FriedMeat.Amount))
	fmt.Println(fmt.Printf("Складники: %s вагою %d", s.FreshCabbage.Name, s.FriedMeat.Amount))
}