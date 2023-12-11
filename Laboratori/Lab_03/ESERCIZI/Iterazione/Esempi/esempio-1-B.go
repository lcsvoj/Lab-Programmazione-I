package main

import "fmt"

func main() {
	var n int = 3
	
	var i int

	i = 0

	for i < n {
	
		fmt.Print("Inizio iterazione numero ", i, ".\n")
	
		// blocco di codice - inizio
		fmt.Println("Istruzioni da eguire al più",n,"volte...")
		// blocco di codice - fine
		
		fmt.Println("Fine iterazione numero ", i, ".\n")
		
		i++
		
	}
	
}
