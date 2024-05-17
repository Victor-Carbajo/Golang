package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprint(x, "\t", y, "\t", z)
	// s := fmt.Sprintf("%v\t%q\t%v\t", x, y, z) é outra solução possível
	fmt.Println(s)
}