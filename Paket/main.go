package main

import (
	"recipes/cabbage"
	"recipes/meat"
	"recipes/shaurma"
)

func main() {
	first := shaurma.Shaurma{
		Name: "Київська",
		FriedMeat: meat.Meat{
			Name: "курка",
			Amount: 700,
		},
		FreshCabbage: cabbage.Cabbage{
			Name: "капуста білокачанна",
			Amount: 500,
		},
	}

	shaurma.EatUp(first)
}