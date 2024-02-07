package main

import "fmt"

func main() {
	var a = [6]int{1, 2, 3, 4, 5, 6}
	stampa(a[:])

	b := a[:]
	b = modifica(b)

	stampa(b)
}

func stampa(sl []int) {
	for _, v := range sl {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func modifica(slIn []int) (slOut []int) {
	slOut = slIn
	for i := range slOut {
		slOut[i] *= 2
	}
	return
}
