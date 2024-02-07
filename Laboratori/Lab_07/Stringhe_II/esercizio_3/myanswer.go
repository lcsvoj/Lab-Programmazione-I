package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	fmt.Print("Inserisci testo (termina con CTRL+D): \n")
	testo := LeggiTestoContinuo()
	parole, lettere := StatisticheParole(testo)
	fmt.Printf("Statistiche: \nNumero parole: %d \nLunghezza media: %.2f\n", int(parole), float64(lettere)/float64(parole))
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

func StatisticheParole(s string) (int, int) {
	var parole, lettere int
	var carattereAnteriore bool = false

	for _, c := range s {
		if unicode.IsLetter(c) {
			if !carattereAnteriore {
				parole++
				lettere++
			} else {
				lettere++
			}
			carattereAnteriore = true
		} else {
			carattereAnteriore = false
		}
	}
	return parole, lettere
}
