package main

import "fmt"

// func main() {
// 	value1 := 1
// 	value2 := value1
// 	value2 += 4
// 	value1 += 2
// 	fmt.Println("Змінні", "Значення1 ", value1, "Значення 2", value2)
// }

func main() {
	value := 9

	ptr1 := &value
	fmt.Println(*ptr1)

	ptr2 := &value
	fmt.Println(ptr2)

	*ptr2 = 7
	fmt.Println(value, *ptr1, *ptr2)
}