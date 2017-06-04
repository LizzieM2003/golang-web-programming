package main

import "fmt"

type gator int

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator.")
}

type flamingo bool

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful.")
}

type swampCreature interface {
	greeting()
}

func bayou(s swampCreature) {
	s.greeting()
}

func main() {
	var g1 gator
	var x int
	g1 = 9
	x = int(g1)
	fmt.Printf("%T\n", g1)
	fmt.Printf("%T\n", x)
	// g1.greeting()
	bayou(g1)
}
