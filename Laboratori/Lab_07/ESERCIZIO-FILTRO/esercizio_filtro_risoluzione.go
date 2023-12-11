package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int

	fmt.Scan(&n)
	numero := strconv.Itoa(n)

	for indice := 0; indice < len(numero); indice++ {
		if indice == 0 {
			fmt.Printf("Output: %s", string(numero[indice]))
			continue
		}

		if numero[indice] != numero[indice-1] {
			fmt.Printf("%s", string(numero[indice]))
		}
	}
	fmt.Println()
}
