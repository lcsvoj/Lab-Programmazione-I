/*
Scrivere un programma che legga da standard input un intero n . Il programma deve stampare a
video la sequenza di numeri che vanno da 1 a n come mostrato nell'Esempio d'esecuzione. In particolare:
- ogni numero divisibile per 3 deve essere rimpiazzato dalla parola "Fizz" ;
- ogni numero divisibile per 5 deve essere rimpiazzato dalla parola "Buzz" ;
- ogni numero divisibile sia per 3 sia per 5 deve essere sostituito da "Fizz Buzz" 
*/

package main

import "fmt"

func main() {
	
	var n int
	fmt.Scan(&n)
	
	for i := 1; i <= n; i++ {
		
		if i%3==0 {
			fmt.Print("Fizz ")
		}
		if i%5==0 {
			fmt.Print("Buzz ")
		}
		if i%3 != 0 && i%5!=0{
			fmt.Print(i, " ")
		}
	}
	fmt.Print("\n")
}

