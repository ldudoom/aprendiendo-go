package main

import (
	"fmt"
	"strings"
)

func main(){
	// Un palindromo es una palabra que puede leerse de igual manera, tando de izquierda a derecha
	// como de derecha a izquierda, por ejemplo OSO
	// Vamos a escribir un codigo utilizando las funciones de la clase strings para detectar palabras
	// que son palindromos
	var palabra string
	fmt.Print("Ingrese una palabra para verificar si es un palindromo:")
	fmt.Scanln(&palabra)
	fmt.Println("Es un palindromo?", esPalindromo(palabra))
}


func esPalindromo(palabra string) bool{
	// OSO
	//fmt.Println(palabra)
	palabra = strings.ToLower(palabra)
	//fmt.Println(palabra)
	
	palabra = strings.ReplaceAll(palabra," ","")
	//fmt.Println(palabra)

	palabraInvertida := reverse(palabra)
	return palabra == palabraInvertida
}

func reverse(cadena string) string{
	aCadena := strings.Split(cadena,"")
	aReverse := make([]string,0)
	for i:=len(aCadena)-1;i>=0;i--{
		aReverse = append(aReverse, aCadena[i])
	}
	return strings.Join(aReverse,"")
}
