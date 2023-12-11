/*

Scrivere un programma che legga da standard input due numeri interi a e b e calcoli il risultato della
divisione a/b . Se b è uguale a 0, il programma stampa Impossibile .

*/


package main

import "fmt"

func main() {
	
	var a, b int
	fmt.Print("Inserisci due numeri: ")
	fmt.Scan(&a, &b)

	if b == 0 {
		fmt.Println("Impossibile")
	} else {
		var risultato float32 = float32(a)/float32(b)		
		fmt.Println("Quoziente =", risultato)
	}

}
