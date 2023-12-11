// Lettura di testo tramite Scanner
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	var testo string
	fmt.Println("Inserisci una riga di testo.")
		
	s := bufio.NewScanner(os.Stdin)  // s è una variabile di tipo bufio.Scanner
	
	if s.Scan() {
		
		testo = s.Text()
		fmt.Print("Letto: ***")
		fmt.Print(testo)
		fmt.Println("+++")
		
	}
}
