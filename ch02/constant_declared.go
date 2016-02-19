package main

import "fmt"

const (
	Pi    = 3.14
	Truth = false
	Big   = 1 << 100
	Small = Big >> 99
)

func main() {
	const Greeting = "ハローワールド"
	fmt.Println(Greeting)
	fmt.Println(Pi)
	fmt.Println(Truth)
}
