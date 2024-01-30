package main

import (
  "bufio"
  "os"
  "strconv"
  "fmt"
  "strings"
  "math"
)

/* Definire la struttura Punto per memorizzare: l' etichetta , l' x e
l' y di un punto sul piano cartesiano. */
type Punto struct {
	etichetta string
	x float64
	y float64
}

func main() {
  tragitto := NuovoTragitto()
  StampaRisultato(tragitto)
}

/* La funzione NuovoTragitto legge da standard input una sequenza di righe di testo nel
formatto definito e restituice la sequenza di istanze del tipo Punto inizializzate con
i valori letti da standard input. */
func NuovoTragitto() (tragitto []Punto) {
	/*  Si assuma che:
	- le righe di testo lette da standard input siano nel formato corretto;
	- la tripla di valori in ogni riga specifichi correttamente un punto sul piano cartesiano;
	- vengano lette da standard input almeno 2 righe di testo.
	*/

  /* Legge da standard input una sequenza di righe di testo nel formato "etichetta;x;y",
  terminando la lettura quando viene letto l'indicatore End-Of-File (EOF) */
  punti := LeggiInput()

  /* Restituisce un valore []Punto nella variabile tragitto in cui è memorizzata la
  sequenza di istanze del tipo Punto inizializzate con i valori letti da standard input.
  L'ordine dei punti all'interno della slice tragitto deve rispettare l'ordine in cui i
  corrispondenti valori sono stati letti da standard input.*/
  for _, punto := range punti {
    datiDelPunto := strings.Split(punto, ";")
    tragitto = append(tragitto, CreaPunto(datiDelPunto))
  }
  return tragitto

}

/* La funzione Distanza riceve in input due punti nei parametri p1 e p2 e restituisce
un valore float64 pari alla distanza euclidea tra di loro */
func Distanza(p1, p2 Punto) float64 {
  return math.Sqrt((math.Pow((p1.x - p2.x), 2) + math.Pow((p1.y - p2.y), 2)))
}

/* La funzione String riceve in input un punto nel parametro p e restituisce una
rappresentazione string di p nel formato "etichetta = (x, y)" */
func String(p Punto) string {
  return fmt.Sprintf("%s = (%.1f, %.1f)", p.etichetta, p.x, p.y)
}

/* La funzione Lunghezza riceve in input una slice di punti nel parametro tragitto e
restituisce un valore float64 pari alla lunghezza del tragitto */
func Lunghezza(tragitto []Punto) float64 {
  var lunghezza float64
  for i := 0; i < len(tragitto)-1; i++ {
    lunghezza += Distanza(tragitto[i], tragitto[i+1])
  }
  return lunghezza
}

/* La funzione LeggiInput legge da standard input una sequenza di righe di testo nel
formato "etichetta;x;y", terminando la lettura quando viene letto l'indicatore EOF */
func LeggiInput() (punti []string) {
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    riga := scanner.Text()
    if riga == "" {
      break
    }
    punti = append(punti, riga)
  }
  return punti
}

func CreaPunto(datiDelPunto []string) (Punto) {
  x, _ := strconv.ParseFloat(datiDelPunto[1], 64)
  y, _ := strconv.ParseFloat(datiDelPunto[2], 64)
  punto := Punto{
    etichetta: datiDelPunto[0],
    x: x,
    y: y,
  }
  return punto
}

func StampaRisultato(tragitto []Punto) {
  var tragittoPercorso float64

  // Calcola e stampa la lunghezza
  lunghezza := Lunghezza(tragitto)
  fmt.Printf("Lunghezza percorso: %.3f\n", lunghezza)

  // Fai il tragitto e, quando si va oltre la mettà delle lunghezza, stampa il prossimo punto
  for i := 0; i < len(tragitto)-1; i++ {
    tragittoPercorso += Distanza(tragitto[i], tragitto[i+1])
    if tragittoPercorso > (lunghezza/2) {
      fmt.Println("Punto oltre metà:", String(tragitto[i+1]))
      break
    }
  }
}
