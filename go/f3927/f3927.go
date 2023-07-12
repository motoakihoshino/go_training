package main

import "fmt"

func main() {
	thenum := 3927
	fmt.Println(thenum)

	for n := 2; n < thenum; n++ {
		if thenum%n == 0 {
			thenum /= n
			fmt.Printf("を%dで割ると、%d\n", n, thenum)
		}
	}
}
