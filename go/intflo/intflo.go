package main

import "fmt"

func main() {
	a := 5 * 2.5
	fmt.Println(a)
	b := 6
	c := float64(b) + 0.2
	fmt.Println(c)

	fmt.Printf("aのデータ型は%T\n", a)
	fmt.Printf("bのデータ型は%T\n", b)
	fmt.Printf("cのデータ型は%T\n", c)
}
