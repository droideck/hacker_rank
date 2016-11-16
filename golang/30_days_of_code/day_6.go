package main

import "fmt"

func main() {
	var input, fword, sword []byte
	var num int

	fmt.Scanf("%d", &num)
	if num < 1 || num > 10 {
		fmt.Println("Input should be >=1 and <=10")
		return
	}
	for i := 0; i < num; i++ {
		fword, sword = nil, nil
		fmt.Scanf("%s", &input)
		if len(input) < 1 || len(input) > 10 {
			fmt.Println("Input length should be >=2 and <=10000")
			continue
		}
		for j := 0; j < len(input); j++ {
			if j%2 == 0 {
				fword = append(fword, input[j])
			} else {
				sword = append(sword, input[j])
			}
		}
		fmt.Println(string(fword), string(sword))
	}
}
