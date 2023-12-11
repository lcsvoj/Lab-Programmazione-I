/* Scrivere un programma che legge da standard input un numero intero n e verifica se il numero è multiplo di 10 */

package main

import "fmt"

func main() {

	var numero int
	fmt.Print("Inserisci numero: ")
	fmt.Scan(&numero)

	if numero%10 == 0 {
		fmt.Println(numero, "è multiplo di 10")
	} else {
		fmt.Println(numero, "non è multiplo di 10")
	}

}
