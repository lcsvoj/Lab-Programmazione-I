package main

import (
	"fmt"
)

func main() {
	/*
	var b [3]int = [3]int{2, 20, 200}
	fmt.Printf("Valore di b:\n%v\n",b)
	fmt.Printf("Tipo e valore di b:\n%#v\n",b)
	*/
	
	//var a[3][2]int 
	//a = [3][2]int{{1, 3}, {10, 30}, {100, 300}}
	
	//fmt.Println(a)
	//fmt.Printf("Tipo e valore di a:\n%#v\n",a)
	
	/*
	for i:=0; i<len(a); i++ {
	    fmt.Printf("i = %d: ",i)
        //fmt.Printf("Tipo e valore di a[%d]:\n%#v\n",i,a[i])
        for j:=0; j<len(a[i]); j++ {
            fmt.Print(a[i][j], " ")
        }
        fmt.Println()
	}
	*/
	
	var m, n int
	m = 4
	n = 3 
    var sm [][]int // sm == nil
    
    /*
    sm = make([][]int, m)
    for i:=0; i<m; i++ {
        sm[i] = make([]int,n)
    }    	
    */
    for i:=0; i<m; i++ {
        sm = append(sm, make([]int, n))
    }
    
    for i:=0; i<len(sm); i++ {
        for j:=0; j<len(sm[i]); j++ {
            sm[i][j] = 100*i+j
        }
    }
    
    
	fmt.Println(sm)
	fmt.Printf("Tipo e valore di sm:\n%#v\n",sm)	

    for i:=0; i<len(sm); i++ {
	    fmt.Printf("i = %d: ",i)
        //fmt.Printf("Tipo e valore di sm[%d]:\n%#v\n",i,sm[i])
        for j:=0; j<len(sm[i]); j++ {
            fmt.Print(sm[i][j], " ")
        }
        fmt.Println()
	}
	
}
