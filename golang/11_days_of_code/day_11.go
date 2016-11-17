package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func makeArray() [6][6]int {
	array := [6][6]int{} // Define a slice with Y cap
	scanner := bufio.NewReader(os.Stdin)

	for i := 0; i < 6; i++ {
		input, _, _ := scanner.ReadLine()
		line := strings.Split(string(input), " ")

		for j := 0; j < 6; j++ {
			num, err := strconv.ParseInt(line[j], 10, 32)
			if err != nil {
				panic(err)
			}
			if num < -9 || num > 9 {
				panic("Number should be >= -9 and <= 9")
			}
			array[i][j] = int(num)
		}
	}

	return array
}

func main() {
	var result int = -63 // Mininum possible result sum
	var sum int

	myArray := makeArray()

	// Count the sum of hourglass
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum = myArray[i][j] + myArray[i][j+1] +
				myArray[i][j+2] + myArray[i+1][j+1] +
				myArray[i+2][j] + myArray[i+2][j+1] + myArray[i+2][j+2]
			if result < sum {
				result = sum
			}
		}
	}
	fmt.Println(result)
}
