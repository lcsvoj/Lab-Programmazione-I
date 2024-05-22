package main

import (
	"fmt"
)

func main() {
	var stringa, nuovaStringa string

	fmt.Print("Inserisci una stringa di testo: ")
	fmt.Scan(&stringa)

	for i := 0; i < len(stringa); i++ {
		nuovaStringa += string(stringa[i]) + " "
	}
	fmt.Println(nuovaStringa)

}
