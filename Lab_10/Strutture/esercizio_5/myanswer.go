/*
Si consideri una rubrica in cui:
- Ogni contatto deve avere un cognome, un nome, un indirizzo ed un numero di telefono.
- Ogni indirizzo deve contenere le seguenti informazioni: via, numero civico, CAP e città.
- Non possono esistere due contatti con lo stesso cognome e lo stesso nome.

Le operazioni sulla rubrica vengano gestite tramite dei comandi letti da standard input. Il programma deve:
1. creare una rubrica vuota;
2. leggere da standard input un testo su più righe (ogni riga è un comando);
3. terminare la lettura quando viene inserito da standard input l'indicatore EOF (CTRL+D);
4. eseguire in sequenza i comando letti eseguendo le funzioni corrispondenti.

Ogni comando è una riga di testo in cui il primo carattere determina il tipo dell'operazione:
	I : inserimento di un contatto			E : cancellazione di un contatto
	S : stampa della rubrica				A : aggiornamento di un contatto

I valori successivi rappresentano i valori da assegnare agli argomenti delle funzioni che effettuano l'operazione richiesta.
I valori sono sempre separati tra di loro dal carattere ";".
	Il comando per l'inserimento di un nuovo contatto ha il seguente formato: "I;nome;cognome;via;numeroCivico;cap;città;telefono"
	Il comando per la cancellazione di un nuovo contatto ha il seguente formato: "E;nome;cognome"

Suggerimento: per separare i valori presenti all'interno di una riga di testo utilizzate la funzione strings.Split().
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	rubrica := NuovaRubrica()
	comandi := LeggiComandi()

	for _, comando := range comandi {
		rubrica = EsegueComando(rubrica, comando)
	}
}

func LeggiComandi() (comandi [][]string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		riga := scanner.Text()
		comandi = append(comandi, strings.Split(riga, ";"))
	}
	return comandi
}

func EsegueComando(r Rubrica, comando []string) Rubrica {

	switch comando[0] {
	case "I":
		numeroCivico, _ := strconv.Atoi(comando[4])
		r = InserisciContatto(r, comando[1], comando[2], comando[3], uint(numeroCivico), comando[5], comando[6], comando[7])
	case "S":
		StampaRubrica(r)
	case "E":
		r = EliminaContatto(r, comando[1], comando[2])
	case "A":
		numeroCivico, _ := strconv.Atoi(comando[4])
		r = AggiornaContatto(r, comando[1], comando[2], comando[3], uint(numeroCivico), comando[5], comando[6], comando[7])
	}
	return r
}

// Restituisce un valore Rubrica rappresentante una rubrica senza alcun contatto inserito.
func NuovaRubrica() Rubrica {
	Rubrica := make([]Contatto, 0)
	return Rubrica
}

// Data la rubrica r, restituisce una nuova rubrica in cui è inserito il nuovo contatto creato con i dati passati come argomento.
// Se la rubrica r contiene già un contatto con identici nome e cognome non avviene nessuna modifica e la rubrica restituita è r stessa.
func InserisciContatto(r Rubrica, nome, cognome string, via string, numero uint, cap, città string, telefono string) Rubrica {
	for _, c := range r {
		if c.nome == nome && c.cognome == cognome {
			return r
		}
	}

	indirizzo := Indirizzo{
		via:          via,
		cap:          cap,
		città:        città,
		numeroCivico: numero,
	}

	contatto := Contatto{
		cognome:   cognome,
		nome:      nome,
		telefono:  telefono,
		indirizzo: indirizzo,
	}

	return append(r, contatto)
}

// Restituisce una rubrica in cui è eliminato il contatto avente nome e cognome uguali a quelli passati in input.
// Se tale contatto non esiste, la rubrica restituita è r stessa.
func EliminaContatto(r Rubrica, nome, cognome string) Rubrica {
	for i, c := range r {
		if c.nome == nome && c.cognome == cognome {
			r = append(r[:i], r[i+1:]...)
			break
		}
	}

	return r
}

// Stampa a video i dettagli del contatto c nel seguente formato: "nome cognome: via numeroCivico, città, cap - Tel. telefono\n"
func StampaContatto(c Contatto) {
	fmt.Printf(
		"%s %s: %s %d, %s, %s - Tel. %s\n",
		c.nome, c.cognome, c.indirizzo.via, c.indirizzo.numeroCivico, c.indirizzo.città, c.indirizzo.cap, c.telefono,
	)
}

/*
Stampa a video tutti i contatti presenti nella rubrica utilizzando per ogni contatto la funzione StampaContatto().
La stampa deve rispettare il seguente formato:

	Rubrica:\n
	[1] - nome cognome: via numeroCivico, città, cap - Tel. telefono\n
	[2] - nome cognome: via numeroCivico, città, cap - Tel. telefono\n
	[3] - nome cognome: via numeroCivico, città, cap - Tel. telefono\n
*/
func StampaRubrica(r Rubrica) {
	fmt.Println("Rubrica:")
	for i, c := range r {
		fmt.Printf("[%d] - ", i+1)
		StampaContatto(c)
	}
}

// Aggiorna i dettagli del contatto identificato da nome e cognome e restituisce la rubrica con il contatto aggiornato.
// Se il contatto non esiste, viene restituita la rubrica r stessa.
func AggiornaContatto(rubrica Rubrica, nome, cognome, via string, numero uint, cap, città, telefono string) Rubrica {
	for i := range rubrica {
		if rubrica[i].nome == nome && rubrica[i].cognome == cognome {
			rubrica[i].indirizzo.via = via
			rubrica[i].indirizzo.numeroCivico = numero
			rubrica[i].indirizzo.cap = cap
			rubrica[i].indirizzo.città = città
			rubrica[i].telefono = telefono
			break
		}
	}
	return rubrica
}

// Indirizzo: Struttura che memorizzi un indirizzo nei seguenti campi:
// via, cap e città di tipo string e numeroCivico di tipo uint.
type Indirizzo struct {
	via          string
	cap          string
	città        string
	numeroCivico uint
}

// Contatto : una struttura che memorizzi i dati di un contatto nei seguenti campi:
// cognome , nome e telefono di tipo string e indirizzo di tipo Indirizzo.
type Contatto struct {
	cognome   string
	nome      string
	telefono  string
	indirizzo Indirizzo
}

// Rubrica : una slice dove ogni elemento è di tipo Contatto.
// Ogni elemento della slice rappresenta un contatto inserito nella rubrica.
// Una slice vuota rappresenta una rubrica vuota.
type Rubrica []Contatto
