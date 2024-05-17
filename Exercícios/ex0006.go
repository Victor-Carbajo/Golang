package main

import "fmt"

type Pharaohmon int

var x Pharaohmon

var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Println("↑ foi x.\n↓ é y!")
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}