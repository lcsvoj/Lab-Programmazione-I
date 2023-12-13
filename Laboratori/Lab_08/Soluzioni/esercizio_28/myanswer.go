package main

import (
	"fmt"
)

func main() {

	var s string
	fmt.Scan(&s)

	numeri := trovaNumeri(s)

	if Decrescente(numeri) {
		fmt.Println("Sequenza nascosta ordinata.")
	} else {
		fmt.Println("Sequenza nascosta non ordinata.")
	}

}

func trovaNumeri(s string) (numeri []int) {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			numeri = append(numeri, int(c)-48)
		}
	}
	return numeri
}

func Decrescente(sl []int) (decrescente bool) {
	for i := 1; i < len(sl); i++ {
		if sl[i] > sl[i-1] {
			return false
		}
	}
	return true
}
