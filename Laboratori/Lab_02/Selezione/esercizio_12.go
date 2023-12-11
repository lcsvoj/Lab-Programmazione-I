/*

Scrivere un unico programma che:
A. legga da standard input un valore intero che specifica il tipo di conversione da effettuare:
	1: secondi (inseriti dall’utente) in ore
	2: secondi inseriti dall’utente in minuti
	3: minuti inseriti dall’utente in ore
	4: minuti inseriti dall’utente in secondi
	5: ore inserite dall’utente in secondi
	6: ore inserite dall’utente in minuti
	7: minuti inseriti dall’utente in giorni e ore
	8: minuti inseriti dall’utente in anni e giorni
	OBS. gestendo l’insertimento di un valore di scelta non compreso tra 1 e 8;
B. legga da standard input un valore reale da convertire;
C. stampi a video il valore convertito.

*/

package main

import "fmt"

func main() {

	var conversione int
	fmt.Print(`
		Scegli la conversione:
		1) secondi -> ore
		2) secondi -> minuti
		3) minuti -> ore
		4) minuti -> secondi
		5) ore -> secondi
		6) ore -> minuti
		7) minuti -> giorni e ore
		8) minuti -> anni e giorni
		: `)
	fmt.Scan(&conversione)	

	var risultato float32

	if (conversione < 1) || (conversione > 8) {
		fmt.Println("Scelta errata")
	} else {

		var valore float32
		fmt.Println("Inserisci il valore da convertire:")
		fmt.Scan(&valore)	

		switch conversione {
		case 1:
			risultato = valore / 3600
			fmt.Println(valore, "secondi corrispondono a", risultato, "ore")
		case 2:
			risultato = valore / 60
			fmt.Println(valore, "secondi corrispondono a", risultato, "minuti")
		case 3:
			risultato = valore * 60
			fmt.Println(valore, "minuti corrispondono a", risultato, "ore")
		case 4:
			risultato = valore * 60
			fmt.Println(valore, "minuti corrispondono a", risultato, "secondi")
		case 5:
			risultato = valore * 3600
			fmt.Println(valore, "ore corrispondono a", risultato, "secondi")
		case 6:
			risultato = valore * 60
			fmt.Println(valore, "ore corrispondono a", risultato, "minuti")
		case 7:
			var giorni, ore float32
			ore = valore / 60
			giorni = ore / 24
			ore = float32(int(ore) % 24)
			fmt.Println(valore, "minuti corrispondono a", int(giorni), "giorni e", ore, "ore")
		case 8:
			var anni, giorni float32
			giorni = valore / (24 * 60)
			anni = giorni / 365
			giorni = float32(int(giorni) % 365)
			fmt.Println(valore, "minuti corrispondono a", int(anni), "anni e", int(giorni), "giorni")
		}
	}
}
