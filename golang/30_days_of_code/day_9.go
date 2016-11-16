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
	if num < 2 || num > 12 {
		fmt.Println("Input should be >=2 and <=12")
		return
	}
	fmt.Println(factorial(num))
}
