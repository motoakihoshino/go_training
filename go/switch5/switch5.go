package main

import "fmt"

func main() {
	commands := []string{"O", "S", "O", "S", "E", "Q"}

	for i := 0; i < len(commands); i++ {
		switch commands[i] {
		case "N":
			fmt.Printf("[%s]新規文章を作成します。\n", commands[i])
		case "O":
			fmt.Printf("[%s]開くファイルを指定してください\n", commands[i])
		case "S":
			fmt.Printf("[%s]ファイルを保存しました\n", commands[i])
		case "Q":
			fmt.Printf("[%s]終了します。よろしいでしょうか\n", commands[i])
		default:
			fmt.Printf("[%s]無効なコマンドです\n", commands[i])
		}
	}
}
