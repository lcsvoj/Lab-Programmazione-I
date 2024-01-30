package main

import "fmt"

func main() {

	if ProprietàSoddisfatta([]rune("1012982101")) {
		fmt.Println("La proprietà testata è soddisfatta.")
	} else {
		fmt.Println("La proprietà testata non è soddisfatta.")
	}

	if ProprietàSoddisfatta([]rune("1212982101")) {
		fmt.Println("La proprietà testata è soddisfatta.")
	} else {
		fmt.Println("La proprietà testata non è soddisfatta.")
	}

}

func ProprietàSoddisfatta(sl []rune) bool {
	fmt.Println(string(sl))
	if len(sl) <= 1 {
		return true
	} else {
		soddisfa := sl[0] == sl[len(sl)-1]
		if soddisfa {
			soddisfa = ProprietàSoddisfatta(sl[1 : len(sl)-1])
		}
		return soddisfa
		// ... equivalente a:
		// soddisfa := sl[0] == sl[len(sl)-1]
		// return soddisfa && ProprietàSoddisfatta(sl[1:len(sl)-1])
	}
}
