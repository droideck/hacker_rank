package main

import "fmt"

func main() {
	var mealCost, tipPercent, taxPercent float32

	fmt.Scanf("%f", &mealCost)
	fmt.Scanf("%f", &tipPercent)
	fmt.Scanf("%f", &taxPercent)

	tip := mealCost * tipPercent / 100
	tax := mealCost * taxPercent / 100
	totalCost := mealCost + tip + tax

	fmt.Printf("The total meal cost is %.f dollars.\n", totalCost)
}
