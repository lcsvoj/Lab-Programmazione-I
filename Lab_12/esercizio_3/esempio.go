package main

import "fmt"

func main() {
	fmt.Println("La somma dei numeri da 0 a 10 è", Somma(10))
}

func Somma(n uint) uint {
	if n == 0 {
		return 0
	} else {
		return n + Somma(n-1)
		// ... equivalente a:
		//var somma_da_0_fino_a_n_meno_1 uint = Somma(n-1)
		//return n + somma_da_0_fino_a_n_meno_1
	}
}
