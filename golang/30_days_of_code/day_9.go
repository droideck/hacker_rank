package main

import "fmt"

func factorial(x uint) uint {
	if x <= 1 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	var num uint
	fmt.Scanf("%d", &num)
	fmt.Println(factorial(num))
}
