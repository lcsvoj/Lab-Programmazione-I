package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var stringheUnite, stringaSingola string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Inserisci una stringa: ")
	for scanner.Scan() {
		fmt.Print("Inserisci una stringa: ")
		stringaSingola = scanner.Text()
		if stringaSingola != "" {
			stringheUnite += stringaSingola + " "
		} else {
			break
		}
	}
	fmt.Println(stringheUnite)
}
