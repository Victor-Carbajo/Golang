package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("%d %T\n", x, x)
	fmt.Printf("%q %T\n", y, y)
	fmt.Printf("%t %T\n", z, z)
	//fmt.Printf("%v\n%q\n%v\n", x, y, z) é outra solução possível
}

// %d: Formata um valor como um número inteiro.
// %f: Formata um valor como um número de ponto flutuante.
// %s: Formata um valor como uma string.
// %q: Imprime uma string com aspas.
// %t: Formata um valor como um booleano.
// %v: Formata o valor na forma padrão para o tipo. Por exemplo, %v com um booleano imprimirá true ou false.
// %T: Formata o tipo do valor.