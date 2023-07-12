package main

import "fmt"

func main() {
	str := "internationalzation"
	strja := "達磨さんが転んだ"

	bt := []byte(str)
	fmt.Println(bt)

	rn := []rune(strja)
	fmt.Println(rn)
}
