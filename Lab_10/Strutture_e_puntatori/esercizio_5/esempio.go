package main

import (
	"fmt"
	"strings"
)

type Persona struct {
	Nome, Cognome string
}

func main() {

	p := &Persona{"Rick", "Sanchez"}

	f(*p)
	fmt.Println(p)
	g(p)
	fmt.Println(p)
}

func f(p Persona) {
	p.Nome, p.Cognome = strings.ToUpper(p.Nome), strings.ToUpper(p.Cognome)
}

func g(p *Persona) {
	p.Nome, p.Cognome = strings.ToUpper(p.Nome), strings.ToUpper(p.Cognome)
}
