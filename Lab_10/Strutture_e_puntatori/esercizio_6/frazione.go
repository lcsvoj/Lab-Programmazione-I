package main

import "fmt"

func main() {
	frazione := NuovaFrazione(34, 18)
	fmt.Println(String(*frazione))
}

type Frazione struct {
	numeratore, denominatore int
}

func NuovaFrazione(numeratore, denominatore int) *Frazione {
	return &Frazione{numeratore, denominatore}
}

func String(f Frazione) string {
	return fmt.Sprintf("%d/%d", f.numeratore, f.denominatore)
}
