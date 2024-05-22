package rettangolo

type Rettangolo struct {
	Base, Altezza float64
}

/* Restituisce una nuova istanza del tipo 'Rettangolo' inizializzata
in base ai valori dei parametri 'base' e 'altezza'. */
func NuovoRettangolo(base, altezza float64) *Rettangolo {
	return &Rettangolo{base, altezza}
}

/* Riceve in input un'instanza del tipo `Rettangolo` nel parametro `r` e
restituisce un valore `float64` pari all'area del rettangolo
rappresentato da `r`. */
func Area(r *Rettangolo) float64 {
	return r.Base * r.Altezza
}
