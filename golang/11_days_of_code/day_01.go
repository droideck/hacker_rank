package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var i uint32 = 4
	var d float32 = 4.0
	var s string = "HackerRank "
	var (
		si uint32
		sd float32
	)

	// Reader object was predefiend in the challenge
	scanner := bufio.NewReader(os.Stdin)
	fmt.Scanf("%d", &si)
	fmt.Scanf("%f", &sd)
	ss, _ := scanner.ReadString('\n')

	fmt.Printf("%d\n", i+si)
	fmt.Printf("%.1f\n", d+sd)
	fmt.Printf("%s\n", s+ss)
}
