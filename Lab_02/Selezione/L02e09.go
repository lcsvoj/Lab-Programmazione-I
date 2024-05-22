package main

import "fmt"

func main() {

	var a, b, c int
	fmt.Print("Inserisci le ampiezze dei due angoli: ")
	fmt.Scan(&a, &b)

	c = 180 - a - b
	if c > 0 {
		fmt.Printf("Ampiezza terzo angolo = %d°", c)
	} else {
		fmt.Println("I due angoli non appartengono ad un triangolo")
	}
}
