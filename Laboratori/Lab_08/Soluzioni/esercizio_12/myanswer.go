package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var moltiplicazione int = 1
	for _, c := range os.Args {
		if n, err := strconv.Atoi(c); err == nil {
			moltiplicazione *= n
		}
	}
	fmt.Printf("Il risultato della moltiplicazione tra i numeri interi è %d\n", moltiplicazione)
}
