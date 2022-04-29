package main

import (
	"fmt"
	"strings"
)

// Closures: Un closure es una funcion que retorna otra funcion, vamos a tener funciones anidadas
// donde la funcion padre va a retornar otra funcion

// Vamos a construir una funcion que imprima n veces una cadena dada
func repeat(n int) func(cadena string) string{
	return func(cadena string) string{
		return strings.Repeat(cadena, n)
	}
}

// funciones anonimas: Son funciones que no tienen un nombre para identificarlas, unicamente son
// declaradas con la palabra freservada func

func main(){

	// funcion anonima
	func(){
		fmt.Println("Hola desde la funcion anonima")
	}() // Se termina con unos parentesis, de esa manera se va a ejecutar automaticamente

	// otra manera es almacenar el resultado de la funcion en una variable
	myFunc := func() {
		fmt.Println("Hola desde la funcion anonima a variable")
	}
	// Para ejecutarla colocamos el nombre de la variable con parentesis como si fuera una funcion
	myFunc()

	// En caso de que la funcion anonima reciva y retorne valores:
	otraFunc := func(nombre string) string{
		return fmt.Sprintf("Hola %s desde la funcion anonima con retorno", nombre)
	}
	fmt.Println(otraFunc("Raul"))

	// Si imprimimos unicamente el nombre de la variable, en este caso otraFunc:
	// fmt.Println(otraFunc)
	// el resultado sera que se imprime el espacio de memoria de esa funcion en formato exadecimal


	// Para probar la funcion closure, lo primero que vamos a hacer es declarar una variable a la que
	// le vamos a enviar el numero entero, es decir vamos a utilizar la funcion padre

	// La variable repeat3 va a almacenar la referencia de memoria de la funcion que retorna repeat()
	repeat3 := repeat(3) 
	// Ahora ya podemos enviar la cadena a la funcion closure y de paso imprimirla en el println
	fmt.Println(repeat3("Hola "))
	fmt.Println(repeat3("Mundo "))

	repeat5 := repeat(5)
	fmt.Println(repeat5("Raul "))
	fmt.Println(repeat5("Chauvin "))
}