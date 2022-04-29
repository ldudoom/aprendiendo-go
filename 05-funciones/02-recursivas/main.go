package main

import "fmt"

func factorial(n int) int{
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main(){
	var numero int
	fmt.Print("Ingrese un numero para calcular su factorial: ")
	fmt.Scanln(&numero)
	fmt.Printf("El factorial de %d es %v\n", numero, factorial(numero))
}