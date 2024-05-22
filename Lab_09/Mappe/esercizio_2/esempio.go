package main

import "fmt"

func main() {
	mappa := make(map[string]int)
	// equivalente a: mappa := map[string]int{}

	mappa["A"] = 10
	mappa["B"] -= 5
	mappa["F"] = 0
	mappa["D"] = mappa["E"] + 5
	if mappa["F"] != 0 {
		fmt.Printf("F è presente con valore %d\n", mappa["F"])
	} else {
		fmt.Print("F non è presente\n")
	}
	if v, ok := mappa["C"]; ok {
		fmt.Printf("C è presente con valore %d\n", v)
	} else {
		fmt.Print("C non è presente\n")
	}
	if v, ok := mappa["B"]; ok {
		fmt.Printf("B è presente con valore %d\n", v)
	} else {
		fmt.Print("B non è presente\n")
	}
	delete(mappa, "B")
	mappa["A"] += 100
	fmt.Println("Elementi in mappa:")
	for k := range mappa {
		fmt.Printf("Chiave: %s - Valore: %d\n", k, mappa[k])
	}
}