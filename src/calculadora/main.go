package main

import "fmt"

func main() {
	fmt.Printf("5 + 5 = %v\n", Soma(5, 5))
}

func Soma(n1 int, n2 int) int {
	return n1 + n2
}