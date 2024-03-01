package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s1, s2 := os.Args[1], os.Args[2]
	fmt.Println(completa(s1, s2))
}

func completa(s1, s2 string) (s3 string) {
	s3 = s1
	for _, c := range s2 {
		if !strings.Contains(s1, string(c)) {
			s3 += string(c)
		}
	}
	return s3
}
