package main

import "fmt"

type person struct {
	fName   string
	lName   string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintf("%s is walking", p.fName)
}

func main() {
	p1 := person{
		fName: "Lizzie",
		lName: "Mendes",
	}
	fmt.Println(p1)
	fmt.Println(p1.fName)
	p1.favFood = []string{"chocolate", "chips", "crisps", "steak"}
	fmt.Println(p1.favFood)
	for _, v := range p1.favFood {
		fmt.Println(v)
	}
	s := p1.walk()
	fmt.Println(s)
}
