package main

import "fmt"

func main() {
	prices := []int{98, 125, 232, 147, 486}
	items := []string{"消しゴム", "ボールペン", "ノート", "ノート", "付箋紙", "バインダー"}
	sum := 0
	for i := 0; i < len(prices); i++ {
		sum += prices[i]
		fmt.Printf("%s:%d円 | 小計%d円\n", items[i], prices[i], sum)
	}

	fmt.Printf("総計:%d円\n", sum)
}
