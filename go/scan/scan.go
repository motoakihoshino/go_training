package main

import "fmt"

func main() {
	var input string

	for {
		fmt.Print("なにか入力してください。''で終了")
		fmt.Scanln(&input)
		if input == "q" {
			break
		}
		fmt.Println(input, "が入力されました")
	}

	fmt.Println("お疲れ様でした。")
}
