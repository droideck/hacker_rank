package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var num int64
	var result int

	fmt.Scanf("%d", &num)
	if num < 1 || num > 1000000 {
		fmt.Println("Input should be >=1 and <=1000000")
		return
	}

	bn := strconv.FormatInt(num, 2)
	bn_list := strings.Split(bn, "0")
	for _, val := range bn_list {
		if result < len(val) {
			result = len(val)
		}
	}
	fmt.Println(result)
}
