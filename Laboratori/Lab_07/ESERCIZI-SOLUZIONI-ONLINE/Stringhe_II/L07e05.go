package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {

	input := LeggiTesto()
	numberStr := ""
	for _, c := range input {
		if unicode.IsNumber(c) {
			numberStr += string(c)
		}
	}

	if numberStr == "" {
		fmt.Println("Nessun numero nascosto")
	} else {
		numberInt, _ := strconv.Atoi(numberStr)
		fmt.Println("Doppio del numero nascosto:", numberInt*2)
	}

}

func LeggiTesto() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Inserisci una riga di testo:")

	// Legge una riga di input fino al newline.
	testoLetto, _ := reader.ReadString('\n')

	// Restituisce il testo letto, escludendo il carattere di newline.
	return testoLetto[:len(testoLetto)-1]
}
