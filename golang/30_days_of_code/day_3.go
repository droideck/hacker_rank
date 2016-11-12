package main

import "fmt"

func main() {
	var num int

	fmt.Scanf("%d", &num)
	switch {
	case (num < 1) && (num > 100):
		fmt.Println("Input should be in range 1 and 100")
	case (num%2 != 0), (num >= 6) && (num <= 20):
		fmt.Println("Weird")
	default:
		fmt.Println("Not Weird")
	}
}
