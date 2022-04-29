package main

import "fmt"

/*
	Las funciones variadicas son funciones que reciben parametros indefinidos, es decir que, podemos
	trabajar con estas funciones cuando no sabemos cuantos parametros va a recibir la funcion

*/

// Vamos a crear una funcion que reciba n numeros para sumarlos
// colocamos el nombre de la variable que va a recibir, el tipo de dato y en el centro colocamos 3 puntos (...)
func sumar(numeros ...int) int{ 
	fmt.Println(numeros)
	suma := 0
	for _, num := range numeros{
		suma += num
	}
	return suma
}

// Vamos a crear una funcion que reciba n valores y que retorne tambien varios valores
// Hay que tener en cuenta que si la funcion recibe mas valores, se debe colocar al final
// la variable que recibira valores indefinidos
func multiplicar(nombre string, numeros ...int) (string, int){
	message := fmt.Sprintf("El producto de %s es: ", nombre)
	producto := 1
	for _, num := range numeros{
		producto *= num
	}
	return message, producto
}

func main(){
	suma := sumar(3,4,5,6,7)
	//fmt.Println("La suma es:",sumar(3,4,5,6,7))
	fmt.Println("La suma es:", suma)
	mensaje, producto := multiplicar("Numeros enteros",3,4,5,6,7)
	//fmt.Println(multiplicar("Numeros enteros",3,4,5,6,7))
	fmt.Println(mensaje, producto)
	// SI no queremos obtener el mensaje para colocar el nuestro personalizado
	_, prod := multiplicar("Numeros enteros",3,4,5,6,7)
	fmt.Println("El producto de los numeros es:", prod)
	
}