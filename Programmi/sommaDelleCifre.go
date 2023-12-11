package main

import "fmt"

func main() {

	var numero int
	for numero < 0 {
		fmt.Print("Inserisci un numero intero non negativo: ")
		fmt.Scan(&numero)
	}

	var somma int
	for numero > 0 {
		somma += numero % 10
		numero /= 10
	}
	fmt.Println("La somma delle cifre è", somma)
}
