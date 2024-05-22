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

	// Leggi il saldo iniziale le operazioni
	saldoIniziale, operazioni := LeggiTesto()

	// Formatta le operazioni
	operazioniFormattate := FormattaOperazioni(operazioni)

	// Crea una slice per guidare l'ordine della mappa
	dateOrdinate := OrdinaDate(operazioniFormattate)

	StampaEstratto(saldoIniziale, operazioniFormattate, dateOrdinate)

}

func LeggiTesto() (saldoIniziale float64, operazioni []string) {

	// Crea un nuovo scanner per leggere riga a riga
	scanner := bufio.NewScanner(os.Stdin)

	// Leggi il saldo iniziale
	scanner.Scan()
	saldoIniziale, _ = strconv.ParseFloat(scanner.Text(), 64)

	for scanner.Scan() {

		// Leggi le operazioni (riga a riga)
		riga := scanner.Text()

		// Ignora le righe vuote
		if riga == "" {
			break
		}

		// Salva ogni operazioni nella slice
		operazioni = append(operazioni, riga)
	}
	return
}

func FormattaOperazioni(operazioni []string) (operazioniFormattate map[string][][]string) {

	operazioniFormattate = make(map[string][][]string)

	for _, operazione := range operazioni {

		// Split: "DATA;TIPO;IMPORTO" -> "DATA", "TIPO", "IMPORTO"
		operazioneSplitatta := strings.Split(operazione, ";")

		// Formatta la data
		dataFormatatta := FormattaData(operazioneSplitatta[0])

		// Formata il tipo di operazione
		switch operazioneSplitatta[1] {
		case "P":
			operazioneSplitatta[1] = "Prelievo"
		case "V":
			operazioneSplitatta[1] = "Versamento"
		}

		// Formatta l'importo della operazione
		importo, _ := strconv.ParseFloat(operazioneSplitatta[2], 64)
		operazioneSplitatta[2] = fmt.Sprintf("%.2f", importo)

		// Salva ["TIPO", "IMPORTO"] associato alla data (gg-mm-aaaa) nella mappa
		operazioniFormattate[dataFormatatta] = append(operazioniFormattate[dataFormatatta], operazioneSplitatta[1:])

	}
	return operazioniFormattate
}

func FormattaData(data string) string {

	dataSplitatta := strings.Split(data, "/")

	// Crea una mappa per armazzenare i dati della data organizzati
	dataFormatatta := make(map[string]string)

	// Caso 1: "aaaa/m/g", "aaaa/m/gg", "aaaa/mm/g" oppure "aaaa/mm/gg"
	if len(dataSplitatta[0]) == 4 {
		dataFormatatta["anno"] = fmt.Sprintf("%04s", dataSplitatta[0])
		dataFormatatta["giorno"] = fmt.Sprintf("%02s", dataSplitatta[2])
	} else {
		// Caso 2: "g/m/aaaa", "gg/m/aaaa", "g/mm/aaaa" oppure "gg/mm/aaaa"
		dataFormatatta["anno"] = fmt.Sprintf("%02s", dataSplitatta[2])
		dataFormatatta["giorno"] = fmt.Sprintf("%02s", dataSplitatta[0])
	}
	// Nei due casi, il mese è sempre nella stessa posizione:
	dataFormatatta["mese"] = fmt.Sprintf("%02s", dataSplitatta[1])

	return fmt.Sprintf("%s-%s-%s", dataFormatatta["giorno"], dataFormatatta["mese"], dataFormatatta["anno"])
}

func OrdinaDate(operazioniFormattate map[string][][]string) (dateOrdinate []string) {
	// Crea una slice per guidare l'ordine della mappa

	var date []string
	for data, _ := range operazioniFormattate {
		date = append(date, data)
	}

	// Inverti le date al formatto "aaaa-mm-gg" per ordinarle con sort.Strings
	var dateInvertiteOrdinate []string
	for _, data := range date {

		// "gg-mm-aaaa" --> "aaaa-mm-gg"
		dataInvertita := fmt.Sprintf("%s-%s-%s", data[6:], data[3:5], data[:2])
		dateInvertiteOrdinate = append(dateInvertiteOrdinate, dataInvertita)
	}

	sort.Strings(dateInvertiteOrdinate)

	// Reverti al formatto anteriore (gg-mm-aaaa)
	for _, data := range dateInvertiteOrdinate {

		// "aaaa-mm-gg" --> "gg-mm-aaaa"
		dataRevertita := fmt.Sprintf("%s-%s-%s", data[8:], data[5:7], data[:4])

		dateOrdinate = append(dateOrdinate, dataRevertita)
	}

	return dateOrdinate
}

func StampaEstratto(saldoIniziale float64, operazioniFormatatte map[string][][]string, dateOrdinate []string) {

	fmt.Printf("\n%-25s%11.2f\n\n", "SALDO INIZIALE: ", saldoIniziale)

	saldoFinale := saldoIniziale
	var tipo string
	var importo float64

	// Ogni data
	for _, data := range dateOrdinate {
		fmt.Println("Data: ", data)
		var saldoGiornaliero float64

		// Ogni operazione in quella data
		for _, operazione := range operazioniFormatatte[data] {
			tipo = operazione[0]
			importo, _ = strconv.ParseFloat(operazione[1], 64)
			fmt.Printf("%-25s%11.2f\n", tipo+":", importo)
			if tipo == "Versamento" {
				saldoGiornaliero += importo
			} else if tipo == "Prelievo" {
				saldoGiornaliero -= importo
			}
		}

		// Al fine della data, atualizza il saldo finale
		saldoFinale += saldoGiornaliero

		fmt.Printf("%-25s%11s\n", "", strings.Repeat("_", 11))

		// Stampa il saldo giornaliero
		fmt.Printf("%-25s%11.2f\n", "SALDO GIORNALIERO:", saldoGiornaliero)
		fmt.Printf("%-25s%11s\n\n", "", strings.Repeat("=", 11))
	}

	fmt.Printf("%-25s%11.2f\n", "SALDO FINALE:", saldoFinale)
	fmt.Printf("%-25s%11s\n\n", "", strings.Repeat("=", 11))

}
