package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var ln int
	fmt.Scanf("%d", &ln)
	scanner := bufio.NewReader(os.Stdin)
	input, _, _ := scanner.ReadLine()
	arr := strings.Split(string(input), " ")

	for i := len(arr) - 1; i >= 0; i-- {
		rs, err := strconv.ParseInt(arr[i], 10, 32)
		if err != nil {
			panic(err)
		}
		if i != 0 {
			fmt.Printf("%d ", rs)
		} else {
			fmt.Printf("%d", rs)
		}
	}
}
