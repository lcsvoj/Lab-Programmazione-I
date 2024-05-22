package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	date := LeggiTesto()

	var dateFormattate []string
	for _, data := range date {
		dateFormattate = append(dateFormattate, Formatta(data))
	}

	sort.Strings(dateFormattate)
	for _, data := range dateFormattate {
		fmt.Println(data)
	}

}

func LeggiTesto() (date []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		if riga == "" {
			return
		}
		date = append(date, riga)
	}
	return
}

func DaInvertire(data string) bool {
	return data[1] == '/' || data[2] == '/'
}

func Inverti(data string) (dataInvertita string) {
	token := ""
	for _, v := range data {
		if v != '/' {
			token += string(v)
		} else {
			dataInvertita = "/" + token + dataInvertita
			token = ""
		}
	}
	dataInvertita = token + dataInvertita

	return
}

func Formatta(data string) string {
	if DaInvertire(data) {
		data = Inverti(data)
	}

	dataRune := []rune(data)

	if len(dataRune) == 8 {
		dataRune = append(dataRune[:4], '-', '0', dataRune[5], '-', '0', dataRune[7])
	} else if len(dataRune) == 9 {
		if dataRune[6] == '/' {
			dataRune = append(dataRune[:4], '-', '0', dataRune[5], '-', dataRune[7], dataRune[8])
		} else {
			dataRune = append(dataRune[:4], '-', dataRune[5], dataRune[6], '-', '0', dataRune[8])
		}
	} else {
		dataRune[4] = '-'
		dataRune[7] = '-'
	}

	return string(dataRune)
}
