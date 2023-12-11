/*

Scrivere un programma che legga da standard input le ampiezze di due angoli di un triangolo e stampi, se
possibile, l'ampiezza del terzo angolo.

*/

package main

import "fmt"

func main() {
	
	var a, b, c int
	fmt.Print("Inserire le ampiezze dei due angoli: ")
	fmt.Scan(&a, &b)
	
	c = 180 - a - b	
	
	if c > 0 {
		fmt.Print("Ampiezza del terzo angolo = ", c, "°\n")
	} else {		
		fmt.Println("I due angoli non appartengono ad un triangolo")
	}

}
