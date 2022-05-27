package main

import "fmt"

func Sum(i, j int) int {
	return i + j
}

func main() {
	fmt.Println("3 + 4 = ", Sum(3,4))
}