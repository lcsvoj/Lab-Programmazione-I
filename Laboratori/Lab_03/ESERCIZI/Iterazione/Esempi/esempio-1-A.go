package main

import "fmt"

func main() {
	var n int = 3
	
	var i int

	for i = 0; i < n; i++ {
		fmt.Print("Inizio iterazione numero ", i, ".\n")
	
		// blocco di codice - inizio
		fmt.Println("Istruzioni da eguire al più",n,"volte...")
		// blocco di codice - fine
		
		fmt.Println("Fine iterazione numero ", i, ".\n")
	}
	
}
