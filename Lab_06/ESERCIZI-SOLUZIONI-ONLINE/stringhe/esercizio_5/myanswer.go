package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var stringaSingola string
	var somma int

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Inserisci un intero: ")
	for scanner.Scan() {
		stringaSingola = scanner.Text()
		if n, err := strconv.Atoi(stringaSingola); err == nil {
			somma += n
			fmt.Print("Inserisci un intero: ")
		} else {
			break
		}
	}
	fmt.Println("Somma:", somma)
}
