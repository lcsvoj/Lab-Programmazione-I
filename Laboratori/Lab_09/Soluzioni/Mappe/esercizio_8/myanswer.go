package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// leggi i dati forniti
	saldoIniziale, operazioni := LeggiTesto()

	// formatta le date
	listaDiDate := creaOrdine(operazioni)

	// stampa il saldo iniziale
	fmt.Printf("%-25s%11.2f\n\n", "SALDO INIZIALE:", saldoIniziale)

	// stampa le operazione e calcola il saldo
	stampaOperazioni(listaDiDate, operazioni)
}

func LeggiTesto() (float64, map[string][]string) {
	fmt.Printf("\n\nStarting the function LeggiTesto()\n\n")
	scanner := bufio.NewScanner(os.Stdin)

	// salva il saldo iniziale
	scanner.Scan()
	saldoIniziale, _ := strconv.ParseFloat(scanner.Text(), 64)

	// salva le operazioni
	operazioni := make(map[string][]string)
	for scanner.Scan() {
		fmt.Printf("\nReading and splitting the line on ';' occurrences\n")
		riga := strings.Split(scanner.Text(), ";") // riga = [data, tipo, importo]
		fmt.Printf("riga = %v\n", riga)
		fmt.Printf("\nAppending the data to `operazioni`, the date has been padronized as %v and used as a key for the value %v\n", padronizzaData(riga[0]), riga[1:])
		operazioni[padronizzaData(riga[0])] = append(operazioni[padronizzaData(riga[0])], riga[1:]...)
		fmt.Printf("operazioni = %v\n", operazioni)
	}
	fmt.Printf("\nFinishing the function, returning variables:\n1. saldoIniziale = %v\n2. operazioni = %v\n\n", saldoIniziale, operazioni)
	return saldoIniziale, operazioni
}

func padronizzaData(data string) string {
	// ritorna una stringa padronizatta YYYY-MM-DD

	dataSlice := strings.Split(data, "/") // dataSlice = [anno, mese, giorno]

	// inverti la data se neccessario
	if len(dataSlice[0]) != 4 {
		dataSlice[0], dataSlice[2] = dataSlice[2], dataSlice[0]
	}

	// mese e giorno devono essere stampati con 2 cifre
	for i := 0; i < 3; i++ {
		if len(dataSlice[i]) == 1 {
			dataSlice[i] = "0" + dataSlice[i]
		}
	}

	// rimonta la data con il formatto giusto
	return fmt.Sprintf(dataSlice[0]) + "-" + dataSlice[1] + "-" + dataSlice[2]
}

func formattaData(data string) string {
	// partendo da una stringa padronizzata YYYY-MM-DD, ritorna una stringa formattata come DD-MM-YYYY

	dataSlice := strings.Split(data, "-") // dataSlice = [anno, mese, giorno]
	return fmt.Sprintf(dataSlice[2]) + "-" + dataSlice[1] + "-" + dataSlice[0]
}

func creaOrdine(operazioni map[string][]string) (ordineDate []string) {
	for data, _ := range operazioni {
		ordineDate = append(ordineDate, data)
		sort.Strings(ordineDate)
	}
	fmt.Println("\nORDINE DATE:", ordineDate)
	return ordineDate
}

func stampaOperazioni(ordineDate []string, operazioni map[string][]string) {

	for _, data := range ordineDate {
		fmt.Println("DATA:", formattaData(data))

		var saldoGiornaliero float64
		for _, operazione := range operazioni[data] {
			fmt.Println("Operazione:", operazione)
			importo := float64(operazione[1])
			if operazione[0] == 'P' {
				fmt.Printf("%-25s%11.2f\n", "Prelievo:", importo)
				saldoGiornaliero -= importo
			} else {
				fmt.Printf("%-25s%11.2f\n", "Versamento:", importo)
				saldoGiornaliero += importo
			}
		}
		fmt.Printf("%-25s%11s\n", "", strings.Repeat("_", 11))
		fmt.Printf("%-25s%11.2f\n", "SALDO GIORNALIERO:", saldoGiornaliero)
		fmt.Printf("%-25s%11s\n\n", "", strings.Repeat("=", 11))
		saldoGiornaliero = 0
	}
}
