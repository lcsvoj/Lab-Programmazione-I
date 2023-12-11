package main

import (
	"fmt"
)

func main() {

	var a, b float64
	fmt.Print("Inserisci due numeri reali: ")
	fmt.Scan(&a, &b)

	radice := -b / a
	fmt.Println("La radice è", radice)
}
