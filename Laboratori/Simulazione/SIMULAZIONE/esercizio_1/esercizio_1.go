package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {

	for i := 1; i < len(os.Args[1:])+1; i++ {
		fmt.Print(TrasformaParola(os.Args[i], i-1))
		fmt.Print(" ")
	}
	fmt.Println()
}

func TrasformaParola(parola string, posizione int) (parolaTrasformata string) {
	/* restituice una stringa in cui i caratteri di parola compaiono alternativamente in maiuscolo
	e in minuscolo:
		- se posizione è pari, il primo carattere in parolaTrasformata è maiuscolo,
		- se posizione è dispari, il primo carattere in parolaTrasformata è minuscolo */

	var anterioreMaiuscola bool
	for i, c := range parola {
		if i == 0 {
			if posizione%2 == 1 {
				parolaTrasformata += string(unicode.ToLower(c))
				anterioreMaiuscola = false
			} else {
				parolaTrasformata += string(unicode.ToUpper(c))
				anterioreMaiuscola = true
			}
		} else {
			if anterioreMaiuscola {
				parolaTrasformata += string(unicode.ToLower(c))
				anterioreMaiuscola = false
			} else {
				parolaTrasformata += string(unicode.ToUpper(c))
				anterioreMaiuscola = true
			}
		}
	}
	return parolaTrasformata
}
