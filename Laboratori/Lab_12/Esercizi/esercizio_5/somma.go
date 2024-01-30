package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Somma valori:", Somma(LeggiInteri()))

}

func LeggiInteri() (valori []int) {
	for _, v := range os.Args[1:] {
		if n, err := strconv.Atoi(v); err == nil {
			valori = append(valori, n)
		}
	}
	return
}

func Somma(valori []int) (somma int) {
	if len(valori) == 0 {
		return 0
	} else {
		return valori[0] + Somma(valori[1:])
	}
}
