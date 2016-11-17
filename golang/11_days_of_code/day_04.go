package main

import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	if initialAge >= 0 {
		p.age = initialAge
	} else {
		p.age = 0
		fmt.Println("Age is not valid, setting age to 0.")
	}
	return p
}

func (p person) amIOld() {
	switch {
	case (p.age < 13):
		fmt.Println("You are young.")
	case (p.age >= 13) && (p.age < 18):
		fmt.Println("You are a teenager.")
	default:
		fmt.Println("You are old.")
	}
}

func (p person) yearPasses() person {
	p.age++
	return p
}

func main() {
	var T, age int

	fmt.Scan(&T)
	if T < 1 || T > 4 {
		fmt.Println("Input should be >=1 and <=4")
		return
	}

	for i := 0; i < T; i++ {
		fmt.Scan(&age)
		if age < -5 || age > 30 {
			fmt.Println("Age should be >=-5 and <=30")
			continue
		}
		p := person{age: age}
		p = p.NewPerson(age)
		p.amIOld()
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}
