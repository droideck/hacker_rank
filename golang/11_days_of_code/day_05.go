package main

import "fmt"

func main() {
	var num int
	fmt.Scanf("%d", &num)
	if num < 2 || num > 20 {
		fmt.Println("Input should be >=2 and <=20")
		return
	}
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
}
