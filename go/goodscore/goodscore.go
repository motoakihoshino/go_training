package main

import "fmt"

func main() {
	scores := [5][3]int{
		{20, 98, 40}, {70, 65, 62}, {35, 39, 38}, {82, 96, 48}, {92, 87, 85},
	}

	for i := 0; i < 5; i++ {
		total := scores[i][0]+scores[i][1]+scores[i][2] > 179
		balance := scores[i][0] > 59 && scores[i][1] > 59 && scores[i][2] > 59
		talent := scores[i][0] > 89 || scores[i][1] > 89 || scores[i][2] > 89
		fmt.Printf("受験者%dさんは：\n", i+1)
		if total {
			fmt.Println("総合点で合格しています")
		}
		if balance {
			fmt.Println("全ての課題で合格しています")
		}
		if !total && !balance && talent {
			fmt.Println("一芸に秀でています")
		}
		if total && balance {
			fmt.Println("おめでとうございます。")
		} else if total || talent {
			fmt.Println("来週、再挑戦できます。")
		} else {
			fmt.Println("来年、また挑戦してください")
		}
		fmt.Println()
	}
}
