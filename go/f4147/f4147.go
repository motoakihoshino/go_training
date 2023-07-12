package main

import "fmt"

func main() {
	thenum := 4147

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
}
