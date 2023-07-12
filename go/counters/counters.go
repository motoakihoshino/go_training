package main

import "fmt"

func counter() func() {
	ctr := 0
	fmt.Println("カウンタを初期化しました。")
	return func() {
		ctr++
		fmt.Println(ctr)
	}
}

func main() {
	countfnc := counter()
	countfnc()
	countfnc()
	countfnc()
}
