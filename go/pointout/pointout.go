package main

import "fmt"

func main() {
	inta := 5
	adinta := &inta
	*adinta = 7

	fmt.Printf("adintaの値は %p\n", adinta)
	fmt.Printf("intaのアドレスは %p\n", &inta)
	fmt.Printf("adintaの内容は%d\n", *adinta)

	fmt.Printf("adintaの値を変更：")
	fmt.Printf("intaの値は %d\n", inta)
}
