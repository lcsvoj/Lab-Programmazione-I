package main

import (
	"fmt"
	"unicode"
)

func main() {

	var parola string
	fmt.Scan(&parola)

	if Ordinata(parola) {
		fmt.Println("Sequenza nascosta ordinata.")
	} else {
		fmt.Println("Sequenza nascosta non ordinata.")
	}
}

func Ordinata(sequenza string) bool {
	precedente := '9'
	for _, v := range sequenza {
		if unicode.IsDigit(v) {
			if v <= precedente {
				precedente = v
			} else {
				return false
			}
		}
	}
	return true
}
