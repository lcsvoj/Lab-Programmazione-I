package main

import "fmt"

func main() {
	conteggio(10)
	fmt.Println("Partenza!")
	fmt.Println()
	conteggioAlternativo(10)
	fmt.Println("Partenza!")
}
func conteggio(n uint) {
	if n != 0 {
		fmt.Println(n)
		conteggio(n - 1)
	}
}
func conteggioAlternativo(n uint) {
	if n != 0 {
		conteggioAlternativo(n - 1)
		fmt.Println(n)
	}
}
