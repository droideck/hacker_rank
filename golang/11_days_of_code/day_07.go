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
	var result string

	fmt.Scanf("%d", &ln)
	if ln < 1 || ln > 1000 {
		fmt.Println("Input should be >=1 and <=1000")
		return
	}
	scanner := bufio.NewReader(os.Stdin)
	input, _, _ := scanner.ReadLine()
	arr := strings.Split(string(input), " ")

	for i := len(arr) - 1; i >= 0; i-- {
		r, err := strconv.ParseInt(arr[i], 10, 32)
		if err != nil {
			panic(err)
		}
		if r < 1 || r > 10000 {
			fmt.Println("Array elements should be >=1 and <=10000 each")
			return
		}

		result += strconv.FormatInt(r, 10) + " "
	}
	fmt.Println(result)
}
