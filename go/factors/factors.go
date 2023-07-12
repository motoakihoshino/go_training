package main

import (
	"fmt"
)

func getfactors(thenum int) {
	fmt.Println(thenum)
	for n := 0; n < thenum; n++ {
		switch {
		case n == 0:
		case n == 1:
		case thenum%n == 0:
			thenum /= n
			fmt.Printf("を%dで割ると、%d\n", n, thenum)
		default:
		}
	}
	fmt.Println()
}

func main() {
	getfactors(2310)
	getfactors(37789)
	getfactors(1256997)
}
