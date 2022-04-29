package main

import "fmt"

func main(){
	a := 20
	b := 10

	// suma
	result := a + b
	fmt.Println("Suma: ",result)

	// Resta
	result = a - b
	fmt.Println("Resta: ",result)

	// Multiplicacion
	result = a * b
	fmt.Println("Producto: ",result)

	// Division
	result = a / b
	fmt.Println("Division: ",result)

	var div float64 = 3.0 / 2.0
	fmt.Println("Division 2: ",div)

	// Modulo de una division
	result = a % b
	fmt.Println("Modulo: ",result)

	result = 3 % 2
	fmt.Println("Modulo: ",result)
}