package main

import "fmt"

type Pharaohmon int

var x Pharaohmon

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}

// "type" criou um tipo chamado "Pharaohmon"  int