/*
Scrivere un programma che legga da standard input il raggio di un cerchio e ne calcoli circonferenza e area
*/

package main

import "fmt"
import "math"

func main() {

	var raggio float32
	fmt.Print("Raggio = ")
	fmt.Scan(&raggio)

	var circonferenza, area float32
	circonferenza = 2*math.Pi*raggio
	area = math.Pi*raggio*raggio

	fmt.Print("Circonferenza = ", circonferenza, "\n", "Area = ", area, "\n")
}
