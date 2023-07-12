package main

import "fmt"

type mynumber struct {
	value int
}

func main() {
	var something int
	var someone string
	var somenumber mynumber
	var somearr [5]int
	var somesl []int

	fmt.Println("somenumber", somenumber)
	fmt.Println("something", something)
	fmt.Println("somearr", somearr)
	fmt.Println("somesl", somesl)

	if someone == "" {
		fmt.Println("someoneはからの文字列")
	} else {
		fmt.Println("someone:", someone)
	}
}
