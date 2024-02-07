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
	if len(insiemeResiduo) == 1 {
		return []string{insiemeResiduo}
	} else {

		sottoinsiemiParziali := Sottoinsiemi(insiemeResiduo[1:])
		for _, sottoinsieme := range sottoinsiemiParziali {
			sottoinsiemi = append(sottoinsiemi, string(insiemeResiduo[0])+sottoinsieme)
		}
		sottoinsiemi = append(sottoinsiemi, sottoinsiemiParziali...)
		sottoinsiemi = append(sottoinsiemi, string(insiemeResiduo[0]))
		
		return		
	}
}
