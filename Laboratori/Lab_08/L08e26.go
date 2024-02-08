package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	date := LeggiTesto()
	dateOrganizzate := make([]string, 0)

	for _, data := range date {
		dateOrganizzate = append(dateOrganizzate, FormattaData(data, "/"))
	}

	sort.Strings(dateOrganizzate)
	for _, data := range dateOrganizzate {
		fmt.Println(data)
	}
}

func LeggiTesto() []string {
	scanner := bufio.NewScanner(os.Stdin)
	righe := make([]string, 0)
	for scanner.Scan() {
		riga := scanner.Text()
		if riga == "" {
			break
		}
		righe = append(righe, riga)
	}
	return righe
}

func FormattaData(data, separatore string) (dataFormatatta string) {
	campi := strings.Split(data, separatore)
	var giorno, mese, anno string
	nuovoSeparatore := "-"

	if len(campi[0]) == 4 {
		anno, mese, giorno = campi[0], campi[1], campi[2]
	} else if len(campi[2]) == 4 {
		giorno, mese, anno = campi[0], campi[1], campi[2]
	}

	if len(mese) == 1 {
		mese = "0" + mese
	}

	if len(giorno) == 1 {
		giorno = "0" + giorno
	}

	return (anno + nuovoSeparatore + mese + nuovoSeparatore + giorno)
}
