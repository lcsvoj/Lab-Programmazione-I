package main

import (
	"fmt"
)

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	fmt.Println(StringheAlternate(s1, s2))
}

func StringheAlternate(s1, s2 string) (risultato string) {
	var nuovaStringa string
	if len(s1) > len(s2) {
		for i := 0; i < len(s1); i++ {
			if i < len(s2) {
				nuovaStringa += string(s1[i])
				nuovaStringa += string(s2[i])
			} else {
				nuovaStringa += string(s1[i])
				nuovaStringa += "-"
			}
		}
		return nuovaStringa
	} else {
		for i := 0; i < len(s2); i++ {
			if i < len(s1) {
				nuovaStringa += string(s1[i])
				nuovaStringa += string(s2[i])
			} else {
				nuovaStringa += "-"
				nuovaStringa += string(s2[i])
			}
		}
		return nuovaStringa
	}
}
