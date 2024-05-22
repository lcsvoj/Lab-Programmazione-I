package main

import (
	"./rettangolo"
	"fmt"
	"os"
	"strconv"
)

func main() {
	base, _ := strconv.ParseFloat(os.Args[1], 64)
	altezza, _ := strconv.ParseFloat(os.Args[2], 64)

	r := rettangolo.NuovoRettangolo(base, altezza)

	fmt.Println(rettangolo.Area(r))
}
