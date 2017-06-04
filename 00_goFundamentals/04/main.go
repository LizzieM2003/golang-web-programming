package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

type transportation interface {
	transportationDevice() string
}

func (t truck) transportationDevice() string {
	return fmt.Sprint("Pickup truck")
}

func (s sedan) transportationDevice() string {
	return fmt.Sprint("Cruising sedan")
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

func main() {
	t := truck{
		vehicle{4, "red"},
		true,
	}
	s := sedan{
		vehicle{2, "black"},
		false,
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.color)
	fmt.Println(s.color)
	report(t)
	report(s)
}
