package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Indirizzo: una struttura che memorizzi un indirizzo nei seguenti campi: via, cap e città di tipo string e numeroCivico di tipo uint ;
type Indirizzo struct {
	Via          string
	Cap          string
	Città        string
	NumeroCivico uint
}

// Contatto: una struttura che memorizzi i dati di un contatto nei seguenti campi: cognome, nome e telefono di tipo string e indirizzo di tipo Indirizzo ;
type Contatto struct {
	Cognome   string
	Nome      string
	Telefono  string
	Indirizzo Indirizzo
}

// Rubrica: una slice dove ogni elemento è di tipo Contatto. Ogni elemento della slice rappresenta un contatto inserito nella rubrica. Una slice vuota rappresenta una rubrica vuota.
type Rubrica []Contatto

func main() {
	r := NuovaRubrica()
	comandi := LeggiComandi()
	for _, comando := range comandi {
		switch comando[0] {
		case "I":
			cognome, nome, via, cap, città, telefono := comando[1], comando[2], comando[3], comando[5], comando[6], comando[7]
			numeroCivico, _ := strconv.Atoi(comando[4])
			r = InserisciContatto(r, cognome, nome, via, uint(numeroCivico), cap, città, telefono)
		case "E":
			cognome, nome := comando[1], comando[2]
			r = EliminaContatto(r, cognome, nome)
		case "S":
			StampaRubrica(r)
		case "A":
			cognome, nome, via, cap, città, telefono := comando[1], comando[2], comando[3], comando[5], comando[6], comando[7]
			numero, _ := strconv.Atoi(comando[4])
			r = AggiornaContatto(r, cognome, nome, via, uint(numero), cap, città, telefono)
		}
	}
}

func LeggiComandi() [][]string {
	comandi := make([][]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		comando := scanner.Text()
		if comando == "" {
			break
		}
		comandi = append(comandi, strings.Split(comando, ";"))
	}
	return comandi
}

// NuovaRubrica restituisce un valore Rubrica rappresentante una rubrica senza alcun contatto inserito;
func NuovaRubrica() Rubrica {
	rubrica := make(Rubrica, 0)
	return rubrica
}

// InserisciContatto, data la rubrica r, restituisce una nuova rubrica in cui è inserito il nuovo contatto creato con i dati passati come argomento.
// Se la rubrica r contiene già un contatto con identici nome e cognome non avviene nessuna modifica e la rubrica restituita è r stessa;
func InserisciContatto(r Rubrica, cognome, nome string, via string, numero uint, cap, città string, telefono string) Rubrica {

	for _, contatto := range r {
		if contatto.Nome == nome && contatto.Cognome == cognome {
			return r
		}
	}

	indirizzo := Indirizzo{
		Via:          via,
		Cap:          cap,
		Città:        città,
		NumeroCivico: numero,
	}

	contatto := Contatto{
		Cognome:   cognome,
		Nome:      nome,
		Telefono:  telefono,
		Indirizzo: indirizzo,
	}

	r = append(r, contatto)
	return r
}

// EliminaContatto restituisce una rubrica in cui è eliminato il contatto avente nome e cognome uguali a quelli passati in input.
// Se tale contatto non esiste, la rubrica restituita è r stessa;
func EliminaContatto(r Rubrica, cognome, nome string) Rubrica {
	rAtualizzata := NuovaRubrica()
	for i, contatto := range r {
		if contatto.Nome == nome && contatto.Cognome == cognome {
			rAtualizzata = append(rAtualizzata, r[:i]...)
			rAtualizzata = append(rAtualizzata, r[i+1:]...)
			return rAtualizzata
		}
	}
	return r
}

// StampaContatto stampa a video i dettagli del contatto c nel seguente formato: "nome cognome: via numeroCivico, città, cap - Tel. telefono\n" ;
func StampaContatto(c Contatto) {
	fmt.Printf("%s %s: %s %d, %s, %s - Tel. %s\n",
		c.Nome, c.Cognome, c.Indirizzo.Via, c.Indirizzo.NumeroCivico, c.Indirizzo.Città, c.Indirizzo.Cap, c.Telefono)
}

// StampaRubrica stampa a video tutti i contatti presenti nella rubrica utilizzando per ogni contatto la funzione StampaContatto() .
// La stampa deve rispettare il seguente formato:
// """
// Rubrica:\n
// [1] - nome cognome: via numeroCivico, città, cap - Tel. telefono\n
// [2] - nome cognome: via numeroCivico, città, cap - Tel. telefono\n
// [3] - nome cognome: via numeroCivico, città, cap - Tel. telefono\n
// """
func StampaRubrica(r Rubrica) {
	fmt.Println("Rubrica:")
	for i, c := range r {
		fmt.Printf("[%d] - ", i+1)
		StampaContatto(c)
	}
}

// AggiornaContatto aggiorna i dettagli del contatto identificato da nome e cognome e restituisce la rubrica con il contatto aggiornato.
// Se il contatto non esiste, viene restituita la rubrica r stessa.
func AggiornaContatto(rubrica Rubrica, cognome, nome string, via string, numero uint, cap, città string, telefono string) Rubrica {
	for _, contatto := range rubrica {
		if contatto.Nome == nome && contatto.Cognome == cognome {
			rAtualizzata := EliminaContatto(rubrica, cognome, nome)
			contattoAggiornato := Contatto{
				Nome:    nome,
				Cognome: cognome,
				Indirizzo: Indirizzo{
					Via:          via,
					NumeroCivico: numero,
					Cap:          cap,
					Città:        città,
				},
				Telefono: telefono,
			}
			rAtualizzata = append(rAtualizzata, contattoAggiornato)
			return rAtualizzata
		}
	}
	return rubrica
}
