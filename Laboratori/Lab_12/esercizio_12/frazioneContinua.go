package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numeri := LeggiValori()
	fmt.Println("Frazione continua:", FrazioneContinua(numeri))
}

func LeggiValori() (valori []int) {
	for _, v := range os.Args[1:] {
		n, _ := strconv.Atoi(v)
		valori = append(valori, n)
	}
	return
}
func FrazioneContinua(numeri []int) float64 {
	if len(numeri) == 0 {
		return 0
	} else if len(numeri) == 1 {
		return float64(numeri[0])
	} else {
		return float64(numeri[0]) + 1/FrazioneContinua(numeri[1:])
	}
}
