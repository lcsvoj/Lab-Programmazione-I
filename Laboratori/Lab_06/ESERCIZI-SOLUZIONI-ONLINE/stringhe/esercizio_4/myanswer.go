package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var stringaSingola string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Inserisci una stringa:")
	for scanner.Scan() {
		stringaSingola = scanner.Text()
		if stringaSingola != "" {
			fmt.Println("Stringa in maiuscolo:", strings.ToUpper(stringaSingola))
			fmt.Println("\nInserisci una stringa:")
		} else {
			break
		}
	}
	fmt.Println("ciao")
}
