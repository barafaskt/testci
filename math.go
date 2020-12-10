package main

import "fmt"

func main() {
	fmt.Println(Soma(10, 10))
}

// Soma função que retorna o valor de dois números inteiros
func Soma(a int, b int) int {
	return a + b
}
