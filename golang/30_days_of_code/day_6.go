package main

import "fmt"

func main() {
	var input, fword, sword []byte
	var num int

	fmt.Scanf("%d", &num)
	for i := 0; i < num; i++ {
		fword, sword = nil, nil
		fmt.Scanf("%s", &input)
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
