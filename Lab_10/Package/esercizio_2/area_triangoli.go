package main

import (
	"./triangolo"
	"fmt"
	"os"
	"strconv"
)

func main() {

	lati := [3]float64{}
	for i, v := range os.Args[1:] {
		lati[i], _ = strconv.ParseFloat(v, 64)
	}

	t, err := triangolo.NuovoTriangolo(lati[0], lati[1], lati[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Perimetro triangolo =", triangolo.Perimetro(t))
	fmt.Println("Area triangolo =", triangolo.Area(t))
}
