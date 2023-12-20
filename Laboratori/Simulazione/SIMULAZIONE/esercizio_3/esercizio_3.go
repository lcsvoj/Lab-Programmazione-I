package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	//"unicode"
)

	// Definire la struttura Punto per memorizzare i dati di un punto sul piano cartesiano
	type Punto struct {
		Etichetta string
		X float64
		Y float64
	}

func main() {

	
	tragitto := NuovoTragitto()
	fmt.Println(tragitto)
	
}


func NuovoTragitto() (tragitto []Punto) {
	/* Legga da standard input una sequenza di righe di testo, ognuna delle quali descrive un punto sul piano cartesiano, terminando la lettura quando viene letto EOF. Ogni riga di testo è una stringa che specifica l'etichetta del punto, l'ascissa e l'ordinata del punto nel seguente formato: "etichetta;x;y" */
	
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
	        var punto Punto
		line := scanner.Text()
		
		var data []string
		var ultimoSeparatore int
		for i, c := range line {
		        
		        if c == ';' {
		                data = append(data, line[ultimoSeparatore+1 : i])
                                ultimoSeparatore = i		      
		        } else if i == len(line) {
		                data = append(data, line[ultimoSeparatore+1:])
		        }
		        
		}
		punto.Etichetta = data[0]
		punto.X, _ = strconv.ParseFloat(data[1], 64)
		punto.Y, _ = strconv.ParseFloat(data[2], 64)
                tragitto = append(tragitto, punto)
                		
	}
	return tragitto
	
	/* Restituice la sequenza di istanze del tipo Punto inizializzate con i valori letti da standard input. L'ordine dei punti all'interno della slice tragitto deve rispettare l'ordine in cui i corrispondenti valori sono stati letti da standard input */
	
}

/*
func Distanza(p1, p2 Punto) float64 {
	// restituisce un valore float64 pari alla distanza euclidea tra i punti rappresentati da p1 e p2
}


func String(p Punto) string {
	// restituisce la rappresentazione string di p nel formato "etichetta = (x, y)"
}


func Lunghezza(tragitto []Punto) float64 {
	// restituisce un valore float64 pari alla lunghezza del tragitto rappresentato da tragitto
}*/
