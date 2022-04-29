package main

import "fmt"

//Variable global
var mensaje string

func funcion1() {
	mensaje = "Hola desde la función uno!"
	fmt.Println(mensaje)
}

func funcion2() {
	mensaje = "Hola desde la función dos!"
	fmt.Println(mensaje)
}

func main() {
	mensaje = "Hola desde la función principal!"
	fmt.Println(mensaje)

	// defer sirve para indicarle al compilador que esta instrucción, en este caso la funcion1 debe ejecutarse al final, 
	// es decir, en este caso, la funcion1 se va a ejecutar luego de que main() se haya ejecutado.
	//defer funcion1() 
	funcion1()
	funcion2()

	fmt.Println(mensaje)
}