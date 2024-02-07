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

	saldoIniziale, operazioni := LeggiTesto()

	var operazioniFormattate []string
	for _, operazione := range operazioni {
		operazioniFormattate = append(operazioniFormattate, FormattaOperazione(operazione))
	}

	operazioniRaggruppate := make(map[string][]string)
	for _, operazione := range operazioniFormattate {
		data := operazione[:10]
		tipoEValore := operazione[11:]
		operazioniRaggruppate[data] = append(operazioniRaggruppate[data], tipoEValore)
	}

	StampaEstrattoConto(saldoIniziale, operazioniRaggruppate)

}

func LeggiTesto() (saldoIniziale float64, operazioni []string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	saldoIniziale, _ = strconv.ParseFloat(scanner.Text(), 64)

	for scanner.Scan() {
		riga := scanner.Text()
		if riga == "" {
			return
		}
		operazioni = append(operazioni, riga)
	}
	return
}

func FormattaOperazione(operazione string) string {
	var indice int
	var v rune
	/* indice in cui inizia (e finisce) la sequenza di byte relativa
	al primo carattere ';' che compare nella stringa operazione;
	operazione[:indice] è la sottostringa relativa alla data
	*/
	for indice, v = range operazione {
		if v == ';' {
			break
		}
	}
	return FormattaData(operazione[:indice], '/', '-') + operazione[indice:]

}

func FormattaData(data string, vecchioSep, nuovoSep rune) string {
	if DaInvertire(data, vecchioSep) {
		data = Inverti(data, vecchioSep)
	}

	dataRune := []rune(data)

	if len(dataRune) == 8 {
		dataRune = append(dataRune[:4], nuovoSep, '0', dataRune[5], '-', '0', dataRune[7])
	} else if len(dataRune) == 9 {
		if dataRune[6] == vecchioSep {
			dataRune = append(dataRune[:4], nuovoSep, '0', dataRune[5], nuovoSep, dataRune[7], dataRune[8])
		} else {
			dataRune = append(dataRune[:4], nuovoSep, dataRune[5], dataRune[6], nuovoSep, '0', dataRune[8])
		}
	} else {
		dataRune[4] = nuovoSep
		dataRune[7] = nuovoSep
	}

	return string(dataRune)
}

func DaInvertire(data string, sep rune) bool {
	return rune(data[1]) == sep || rune(data[2]) == sep
}

func Inverti(data string, sep rune) (dataInvertita string) {
	token := ""
	for _, v := range data {
		if v != sep {
			token += string(v)
		} else {
			dataInvertita = string(sep) + token + dataInvertita
			token = ""
		}
	}
	dataInvertita = token + dataInvertita

	return
}

func StampaEstrattoConto(saldoIniziale float64, operazioniRaggruppate map[string][]string) {

	fmt.Printf("%-25s%11.2f\n\n", strings.ToUpper("Saldo iniziale:"), saldoIniziale)

	saldoFinale := saldoIniziale

	var date []string
	for data := range operazioniRaggruppate {
		date = append(date, data)
	}

	sort.Strings(date)

	for _, data := range date {
		fmt.Printf("DATA: %s\n", Inverti(data, '-'))
		var saldoGiornaliero float64
		for _, tipoEValore := range operazioniRaggruppate[data] {
			tipo := tipoEValore[0]
			importo, _ := strconv.ParseFloat(tipoEValore[2:], 64)
			if tipo == 'V' {
				saldoGiornaliero += importo
				fmt.Printf("%-25s%11.2f\n", "Versamento:", importo)
			} else {
				saldoGiornaliero -= importo
				fmt.Printf("%-25s%11.2f\n", "Prelievo:", importo)
			}
		}
		fmt.Printf("%-25s%11s\n", "", strings.Repeat("_", 11))
		fmt.Printf("%-25s%11.2f\n", strings.ToUpper("Saldo giornaliero:"), saldoGiornaliero)
		fmt.Printf("%-25s%11s\n\n", "", strings.Repeat("=", 11))
		saldoFinale += saldoGiornaliero
	}

	fmt.Printf("%-25s%11.2f\n", strings.ToUpper("Saldo finale:"), saldoFinale)
	fmt.Printf("%-25s%11s\n\n", "", strings.Repeat("=", 11))

}
