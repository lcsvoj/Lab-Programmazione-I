package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var saldoIniziale float64

type Operazione struct {
	data    string
	tipo    string
	importo float64
}

func main() {
	// Leggi le righe con delle operazioni dal file
	righe := LeggiOperazioni()

	// Orgenizza le operazioni in una mappa
	operazioni := OrganizzaOperazioni(righe)

	// Crea una slice guida per la stampa, d'accordo con le date ordinate
	dateOrdinate := OrdinaDate(operazioni)

	// Stampa lo estratto
	StampaEstratto(operazioni, dateOrdinate)
}

func LeggiOperazioni() (righe []string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		riga := scanner.Text()
		if riga == "" {
			break
		}

		// La prima riga del file è una stringa nel formato "IMPORTO",
		// un valore reale che specifica il saldo iniziale di un conto corrente.
		if !strings.Contains(riga, ";") {
			saldoIniziale, _ = strconv.ParseFloat(riga, 64)
			continue
		}

		// Ogni riga successiva alla prima è una stringa nel formato "DATA;TIPO;IMPORTO"
		// e specifica un'operazione bancaria effettuata sul conto corrente
		righe = append(righe, riga)
	}
	return righe
}

func FormattaOperazione(riga string) (operazione Operazione) {
	// Ogni riga è una stringa nel formato "DATA;TIPO;IMPORTO"
	campi := strings.Split(riga, ";")

	// campi[0] -> data
	operazione.data = FormattaData(campi[0], "s")

	// campi[1] -> tipo
	switch campi[1] {
	case "P":
		operazione.tipo = "Prelievo"
	case "V":
		operazione.tipo = "Versamento"
	}

	// campi[2] -> importo
	operazione.importo, _ = strconv.ParseFloat(campi[2], 64)

	return operazione
}

func OrganizzaOperazioni(righe []string) (operazioniMap map[string][]Operazione) {
	operazioniMap = make(map[string][]Operazione, 0)
	for _, riga := range righe {
		operazione := FormattaOperazione(riga)
		operazioniMap[FormattaData(operazione.data, "o")] = append(operazioniMap[FormattaData(operazione.data, "o")], operazione)
	}
	return operazioniMap
}

// La funzione FormattaData riceve due input: una stringa con la data in formati specificati (aaaa/m/g, aaaa/m/gg,
// aaaa/mm/g, aaaa/mm/gg, g/m/aaaa, gg/m/aaaa, g/mm/aaaa, gg/mm/aaaa) e un marcatore che definisce il formato di
// output ("s" per formatto di stampa (gg-mm-aaaa), "o" per formatto di ordinamento (aaaammgg))
func FormattaData(s string, formatto string) (dataFormatatta string) {
	// output desiderato: gg-mm-aaaa

	// troviamo il separatore usato nel input
	var separatore string
	for _, c := range s {
		if !unicode.IsNumber(c) {
			separatore = string(c)
		}
	}

	campi := strings.Split(s, separatore)
	var giorno, mese, anno string

	// Troviamo l'anno con base nella sua lunghezza
	for i := range campi {
		if len(campi[i]) == 4 {
			anno = campi[i]

			// Con base nella posizione del anno, sappiamo dov'è il giorno
			if i == 0 {
				giorno = campi[2]
			} else if i == 2 {
				giorno = campi[0]
			}
		}
	}
	// Il mese sarà sempre in campi[1]
	mese = campi[1]

	// Mettiamo un zero nei giorni o mesi scritti con 1 cifra
	if len(giorno) == 1 {
		giorno = "0" + giorno
	}

	if len(mese) == 1 {
		mese = "0" + mese
	}

	// Restituimo la string formatatta d'accordo con il marcatore specificato
	if formatto == "s" {
		// gg-mm-aaaa
		dataFormatatta = giorno + "-" + mese + "-" + anno
	} else if formatto == "o" {
		dataFormatatta = anno + "-" + mese + "-" + giorno
	}
	return dataFormatatta
}

// La funzione OrdinaDate restituice una string con le date in ordine crescente
func OrdinaDate(operazioni map[string][]Operazione) (dateOrdinate []string) {
	dateOrdinate = make([]string, 0)
	for data, _ := range operazioni {
		// Estrai la data per ordinazione
		dateOrdinate = append(dateOrdinate, data)
	}
	sort.Strings(dateOrdinate)
	return dateOrdinate
}

func StampaEstratto(operazioni map[string][]Operazione, dateOrdinate []string) {
	// Stampa il saldo iniziale
	fmt.Printf("SALDO INIZIALE: %20.2f\n", saldoIniziale)

	saldoFinale := saldoIniziale

	// Stampa ogni data e le operazione contenute
	for _, data := range dateOrdinate {

		fmt.Printf("\nDATA: %s\n\n", FormattaData(data, "s"))

		saldoGiornaliero := 0.0
		for _, operazione := range operazioni[data] {
			fmt.Printf("%-25s%11.2f\n", operazione.tipo, operazione.importo)
			if operazione.tipo == "Versamento" {
				saldoGiornaliero += operazione.importo
			} else {
				saldoGiornaliero -= operazione.importo
			}
		}
		fmt.Printf("%-25s%11s\n", "", strings.Repeat("_", 10))
		fmt.Printf("%-25s%11.2f\n", "SALDO GIORNALIERO", saldoGiornaliero)
		fmt.Printf("%-25s%11s\n\n", "", strings.Repeat("=", 10))

		saldoFinale += saldoGiornaliero
	}
	fmt.Printf("\n%-25s%11.2f\n", "SALDO FINALE:", saldoFinale)
	fmt.Printf("%-25s%11s\n\n", "", strings.Repeat("=", 10))

}
