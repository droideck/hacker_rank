package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var phoneBook map[string]string
	var s []byte
	var sarr []string
	var num int

	phoneBook = make(map[string]string)
	scaner := bufio.NewReader(os.Stdin)

	fmt.Scanf("%d", &num)
	for i := 0; i < num; i++ {
		s, _, _ = scaner.ReadLine()
		sarr = strings.Split(string(s), " ")
		phoneBook[sarr[0]] = sarr[1]
	}
	for {
		s, _, _ = scaner.ReadLine()
		if s != nil {
			if _, ok := phoneBook[string(s)]; ok {
				fmt.Printf("%s=%s\n", string(s), phoneBook[string(s)])
			} else {
				fmt.Println("Not found")
			}
		} else {
			break
		}
	}
}
