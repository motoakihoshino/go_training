package main

import "fmt"

func counter() func() {
	ctr := 0
	fmt.Println("カウンタを初期化しました。")
	fmt.Printf("カウンタのアドレスは%p\n", &ctr)
	return func() {
		ctr++
		fmt.Printf("カウンタの値は%d、", ctr)
		fmt.Printf("カウンタのアドレスは%p\n", &ctr)
	}
}

func main() {
	countfnc := counter()
	countfnc()
	countfnc()
	countfnc()
}
