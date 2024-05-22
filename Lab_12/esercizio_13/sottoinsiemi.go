package main

import (
	"fmt"
)

func main() {

	insiemeResiduo := ""

	fmt.Scan(&insiemeResiduo)

	sottoinsiemi := Sottoinsiemi(insiemeResiduo)
	fmt.Printf("%d sottoinsiemi: %v\n", len(sottoinsiemi), sottoinsiemi)

}

func Sottoinsiemi(insiemeResiduo string) (sottoinsiemi []string) {
	if len(insiemeResiduo) == 0 {
		return
	} else {
		sottoinsiemi = Sottoinsiemi(insiemeResiduo[1:])
		for _, sottoinsieme := range sottoinsiemi {
			sottoinsiemi = append(sottoinsiemi, string(insiemeResiduo[0])+sottoinsieme)
		}
		sottoinsiemi = append(sottoinsiemi, string(insiemeResiduo[0]))
		return
	}
}
