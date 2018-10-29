package main

import "fmt"

func main() {
	next := AStarNextStep()
	i := next()
	fmt.Println("s:", i)
	i = next()
	fmt.Println("s:", i)
	i = next()
	fmt.Println("s:", i)
	i = next()
	fmt.Println("s:", i)
	i = next()
	fmt.Println("s:", i)
}