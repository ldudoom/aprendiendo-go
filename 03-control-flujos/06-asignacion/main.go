package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Por favor ingrese un numero entero: ")
	fmt.Scanln(&numero)

	// Operadores de asignacion
	a := numero
	// Suma en asignacion
	a += 2 // a = a + 2
	fmt.Println("Suma: ", a)

	// Resta en asignacion
	a -= 2 // a = a - 2
	fmt.Println("Resta: ", a)

	// Multiplicacion en asignacion
	a *= 2 // a = a * 2
	fmt.Println("Producto: ", a)

	// Division en asignacion
	a /= 2 // a = a / 2
	fmt.Println("Division: ", a)

	// Modulo en Asignacion
	a %= 2 // a = a % 2
	fmt.Println("Modulo: ", a)

}