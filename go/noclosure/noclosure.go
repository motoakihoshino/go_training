package main

import "fmt"

func counter(adctr *int) {
	*adctr++
}

func freecounter(adctr *int, add int) {
	*adctr += add
}

func main() {
	ctr := 0
	fctr := 10

	for i := 0; i < 5; i++ {
		counter(&ctr)
		fmt.Printf("カウンタの値は%d\n", ctr)

		freecounter(&fctr, ctr)
		fmt.Printf("フリーカウンターの値は%d\n", fctr)
	}

}
