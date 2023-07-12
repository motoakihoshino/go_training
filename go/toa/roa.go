package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 9 * 9
	spfa := fmt.Sprintf("9かける9は%d", 9*9)
	fmt.Println(spfa)

	astr := "9の2乗も" + strconv.Itoa(a)
	fmt.Println(astr)

	b := 3.14159

	sbfb := fmt.Sprintf("円周率はだいたい%.5f", b)
	fmt.Println(sbfb)

	bstr := strconv.FormatFloat(b, 'f', 2, 64)
	fmt.Println("もっと大雑把には" + bstr)
}
