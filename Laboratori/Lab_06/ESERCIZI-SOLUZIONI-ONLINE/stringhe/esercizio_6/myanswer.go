package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var stringaSingola string
	var ultimoCognome = " "

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Inserisci un cognome: ")
	for scanner.Scan() {
		stringaSingola = scanner.Text()
		if stringaSingola != "" {
			if stringaSingola[0] > ultimoCognome[0] {
				ultimoCognome = stringaSingola
			}
			fmt.Print("Inserisci un cognome: ")
		} else {
			break
		}
	}
	fmt.Println(ultimoCognome)
}
