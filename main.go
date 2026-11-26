package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func run() {
	fmt.Println(Add(2, 3))
}

func main() {
	run()
}
