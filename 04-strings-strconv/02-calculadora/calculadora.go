package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	var expresion string
	//var result int

	fmt.Print("=> ")
	fmt.Scanln(&expresion)
	resultado := sumar(expresion)
	if(resultado != -1){
		fmt.Println("La suma de los valores es:",resultado)
	}
	
}

func sumar(expresion string) int{
	valores := strings.Split(expresion, "+")
	var result int

	// fmt.Println(valores)
	// fmt.Printf("%T\n",valores)

	// Vamos a recorrer el arreglo de numeros que actualmente son caracteres utilizando
	// un bucle for haciendolo trabajar como un foreach y sin almacenar el indice, por eso el caracter "_"
	result = 0
	for _,valor := range valores {
		// Ahora vamos a utilizar el metodo Atoi de strconv para convertir el caracter en un valor entero
		// devuelve como segundo parametro un nil que de momento no lo vamos a tomar en cuenta
		// por ese motivo colocamos el "_"
		//num, _ := strconv.Atoi(valor)

		// Ahora vamos a realizar un poco de manejo de errores, para eso ya vamos a tomar el cuenta el segundo
		// valor devuelto por el metodo Atoi
		num, err := strconv.Atoi(valor)
		//fmt.Println(num, err)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error: Por favor ingrese unicamente valores enteros y unicamente se acepta el operador '+'")
			result = -1
			break
		} else {
			result += num
		}
	}

	return result
}