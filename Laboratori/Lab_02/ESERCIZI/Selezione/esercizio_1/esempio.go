/* Cercate di capire qual è l'output di questo programma prima di eseguirlo. */
package main

import "fmt"

func main() {
	var (
		a, b int = 10, 20
		c int = 30
	)
	if a > b {
		a = b
	} else {
		b = a
	}
	c = c + b + a
	fmt.Println(a, b, c)
}
