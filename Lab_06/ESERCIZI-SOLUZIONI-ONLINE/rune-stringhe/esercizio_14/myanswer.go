package main

import "fmt"

func main() {
	var parola string
	fmt.Print("Inserisci una parola: ")
	fmt.Scan(&parola)

	if ÈPalindroma(parola) {
		fmt.Println("Palindroma")
	} else {
		fmt.Println("Non palindroma")

	}
}

func ÈPalindroma(parola string) bool {
	var parolaInversa string
	for _, c := range parola {
		parolaInversa = string(c) + parolaInversa
	}
	return parola == parolaInversa
}
