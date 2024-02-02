package main

import (
	"fmt"
)

func main() {

	var distanzaKm float64
	fmt.Print("Distanza (Km) = ")
	fmt.Scan(&distanzaKm)

	var distanzaMi float64 = distanzaKm * 0.62137
	fmt.Println("Distanza (mi) =", distanzaMi)
}
