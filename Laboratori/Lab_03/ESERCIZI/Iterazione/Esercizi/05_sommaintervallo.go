/* Scrivere un programma che legga da standard input due numeri interi a e b e stampi a video la
somma dei numeri pari compresi tra a e b (estremi esclusi). */

package main

import "fmt"

func main() {

	var a, b int
	var somma int

	fmt.Scan(&a, &b)

	for i := a; i <= b; i++ {
		if i%2 == 0 {
			somma = somma + i
		}
	}
	fmt.Println("Somma =", somma)

}
