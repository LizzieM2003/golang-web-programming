package main

import "fmt"

func main() {
	s := "i'm sorry dave i can't do that"
	fmt.Println(s)
	xs := []byte(s)
	fmt.Println(xs)
	fmt.Println(xs, string(xs))
	fmt.Println(s[:14])
	fmt.Println(s[10:22])
	fmt.Println(s[17:])
	for _, v := range s {
		fmt.Println(string(v))
	}
}
