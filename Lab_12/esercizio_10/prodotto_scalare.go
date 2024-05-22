package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	a := LeggiSequenza(n)
	b := LeggiSequenza(n)

	prodotto := Moltiplica(a, b)
	fmt.Println("Prodotto =", prodotto)
}

func LeggiSequenza(n int) (sequenza []int) {
	if n == 0 {
		return
	} else {
		var valore int
		fmt.Scan(&valore)
		return append([]int{valore}, LeggiSequenza(n-1)...)
	}
}
func Moltiplica(a, b []int) (prodotto int) {
	if len(a) == 0 {
		return
	} else {
		return a[0]*b[0] + Moltiplica(a[1:], b[1:])
	}
}
