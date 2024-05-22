package main

import (
	"os"
	"unicode"
  "fmt"
)

func main() {
	parole := os.Args[1:]

  for i, parola := range parole {
    fmt.Print(TrasformaParola(parola, i))
		fmt.Print(" ")
    }
}

func TrasformaParola(parola string, posizione int) (parolaTrasformata string){
		// le parole in posizione pari nella sequenza letta vengono ristampate incominciando con un carattere maiuscolo
		if posizione%2 == 0 {
			for i, lettera := range parola {
				if i%2 == 0 {
					parolaTrasformata += string(unicode.ToUpper(lettera))
				} else {
					parolaTrasformata += string(unicode.ToLower(lettera))
				}
			}
		}

    // le parole in posizione dispari nella sequenza letta vengono ristampate incominciando con un carattere minuscolo
    if posizione%2 == 1 {
			for i, lettera := range parola {
				if i%2 == 0 {
					parolaTrasformata += string(unicode.ToLower(lettera))
				} else {
					parolaTrasformata += string(unicode.ToUpper(lettera))
				}
			}
		}
    return parolaTrasformata
}
