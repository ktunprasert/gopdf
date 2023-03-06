package main

import "fmt"

func mains() {
	me := 32
	fmt.Println("first 15")
	for i, n := me, 1; i > 15; i -= 15 {
		fmt.Println("next 15", i, n, n*15 +1)
		n++
	}
}
