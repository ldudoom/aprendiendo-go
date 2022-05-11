package main

import "fmt"

// Funcion que recibe una copia de la variable
func modificarValor(cadena string){
	fmt.Printf("%p\n", &cadena) // Imprimimos la referencia de memoria
	cadena = "Hola desde la funcion" // le asignamos un nuevo valor
}

// Funcion que recibe la referencia de la memoria de la variable
func modificarValor2(cadena *string){
	fmt.Printf("%p\n", cadena) // Imprimimos la referencia de memoria
	*cadena = "Hola desde la funcion" // le asignamos un nuevo valor
}

func main(){
	cadena := "Hola desde main"

	fmt.Printf("%p\n", &cadena)
	fmt.Println("Antes de la funcion:", cadena)
	// En este caso enviamos una copia de la variable, no la referencia de memoria
	modificarValor(cadena)
	fmt.Println("Despues de la funcion:", cadena)

	
	
	fmt.Printf("%p\n", &cadena)
	fmt.Println("Antes de la funcion2:", cadena)
	// En este caso enviamos la referencia de memoria de la variable, no una copia
	modificarValor2(&cadena)
	fmt.Println("Despues de la funcion2:", cadena)
}