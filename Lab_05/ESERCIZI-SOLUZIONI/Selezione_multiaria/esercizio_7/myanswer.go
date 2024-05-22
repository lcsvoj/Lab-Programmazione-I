package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c float64
	fmt.Print("Inserisci tre numeri reali: ")
	fmt.Scan(&a, &b, &c)

	delta := b*b - 4*a*c
	switch {
	case delta == 0:
		fmt.Print("Due radici reali coincidenti: ", (-b+math.Sqrt(float64(delta)))/2*a)

	case delta > 0:
		fmt.Print("Due radici reali: ", (-b+math.Sqrt(float64(delta)))/2*a, (-b-math.Sqrt(float64(delta)))/2*a)
	case delta < 0:
		var parteReale, parteImaginaria float64
		parteReale = -b / 2 * a
		parteImaginaria = math.Sqrt(-delta) / 2 * a
		fmt.Print("Due radici complesse: ", complex(parteReale, parteImaginaria), complex(parteReale, parteImaginaria))
	}
	fmt.Println()

}
