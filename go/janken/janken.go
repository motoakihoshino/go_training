package main

import "fmt"

func main() {
	pattern := [3]string{"グー", "チョキ", "パー"}

	fmt.Printf("***じゃんけん勝負リスト***\n")

	for me := 0; me < 3; me++ {
		fmt.Printf("わたしが%sのとき：\n", pattern[me])
		for you := 0; you < 3; you++ {
			socre := (3 + me - you) % 3
			if socre == 2 {
				fmt.Printf("あなたが%sならわたしの勝ち\n", pattern[you])
			} else if socre == 1 {
				fmt.Printf("あなたが%sならわたしの負け\n", pattern[you])
			} else {
				fmt.Printf("あなたが%sならあいこ\n", pattern[you])
			}
		}
		fmt.Println()
	}
}
