package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Print("Inserisci testo (termina con CTRL+D): \n")
	fmt.Print("Testo letto:\n", LeggiTestoContinuo())
}

func LeggiTestoContinuo() string {
	in := bufio.NewReader(os.Stdin)
	var result string
	for {
		line, err := in.ReadString('\n')
		if err == io.EOF {
			break
		}
		result += line
	}
	return result
}
