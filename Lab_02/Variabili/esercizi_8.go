/*
Scrivere un programma che legga da standard input una distanza in Km ed effettui la conversione di tale
distanza in miglia (1 Km = 0.62137 mi).
*/

package main

import "fmt"

func main() {

	var distanza_km, distanza_mi float32
	fmt.Print("Distanza (Km) = ")
	fmt.Scan(&distanza_km)
	distanza_mi = distanza_km * 0.62137
	fmt.Println("Distanza (mi) = ", distanza_mi)
}
