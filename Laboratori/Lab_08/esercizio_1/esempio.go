package main

import (
	"fmt"
	"strconv"
)

func main() {

	var numeroIntero int = 7

	fmt.Printf("Tipo: %T\n", numeroIntero)                               // int
	fmt.Printf("- formato default: %v\n", numeroIntero)                  // 7
	fmt.Printf("- formato base 10: %d\n", numeroIntero)                  // 7
	fmt.Printf("- formato base 10 (larghezza = 6): %6d\n", numeroIntero) //      7
	fmt.Printf("- formato base 2: %b\n", numeroIntero)                   // 111

	fmt.Println()

	var numeroReale float64 = 14.5674
	fmt.Printf("Tipo: %T\n", numeroReale)                                       // float64
	fmt.Printf("- formato default: %v\n", numeroReale)                          // 14.5674
	fmt.Printf("- formato con decimali: %f\n", numeroReale)                     // 14.567400
	fmt.Printf("- formato con numero fissato di decimali: %.2f\n", numeroReale) // 14.57
	fmt.Printf("- formato esponenziale: %e\n", numeroReale)                     // 1.456740e+01

	fmt.Println()

	var valoreLogico bool = false
	fmt.Printf("Tipo: %T\n", valoreLogico)                 // bool
	fmt.Printf("- formato default: %v\n", valoreLogico)    // false
	fmt.Printf("- formato true/false: %t\n", valoreLogico) // false

	fmt.Println()

	var carattere rune = 'A'
	fmt.Printf("Tipo: %T\n", carattere)                // int32
	fmt.Printf("- formato default: %v\n", carattere)   // 65
	fmt.Printf("- formato carattere: %c\n", carattere) // A
	fmt.Printf("- formato intero: %d\n", carattere)    // 65
	fmt.Printf("- formato unicode: %U\n", carattere)   // U+0041

	fmt.Println()

	var stringaTesto string = "Hello\tworld!"
	fmt.Printf("Tipo: %T\n", stringaTesto)              // string
	fmt.Printf("- formato default: %v\n", stringaTesto) // Hello	world!
	fmt.Printf("- formato stringa: %s\n", stringaTesto) // Hello	world!

	fmt.Println()

	var posizioniDecimali = 9
	fmt.Printf("Printf con formato creato dinamicamente %."+strconv.Itoa(posizioniDecimali)+"f\n", 10.458)
}
