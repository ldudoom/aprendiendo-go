package main

import "fmt"

func main(){
	/*
	Las variables en GO se definen en primer lugar con la palabra reservada "var"
	GO es un lenguaje tipado, por lo que se debe definir tambien el tipo de dato
	*/
	var nombre1 string
	nombre1 = "Raul"
	
	// Se define la variable y se asigna un valor
	var nombre2 string = "Alejandro"

	// Podemos simplificar la definicion de una variable utilizando el operador :=, de esta manera
	// se define la variable, se le asigna un valor inicial, y no necesitamos definir el tipo de dato
	// ya que la variable va a tomar el tipo de dato de acuerdo al valor que le hayamos asignado
	// Si le asignamos un valor entero, la variablr se definira como int, si le asignamos un valor decimal
	// la variable tomara el tipo de dato float etc
	edad := 40

	fmt.Println("Nombre: ",nombre1, nombre2, "Edad: ", edad)

	// Si no asignamos ningun valor a las variables, hay que tener cuidado ya que estas variables no van a
	// tomar un valor null o undefined, sino que siempre van a tener un valor por defecto
	// En el caso de tipos de dato numericos el valor por defecto sera 0
	// En el caso de cadenas sera un valor vacio
	// En el caso de una variable de tipo boll va a tomar por defecto el valor FALSE
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a,b,c,d)

	// Cuando definimos una constante, como su nombre lo indica, esta constante debera tener un valor
	// inicial por defecto, la constante tomara el tipo de dato de acuerdo al valor que le asignemos
	// Y no podra cambiarse su valor posterior a su definicion
	const pi = 3.141592
	fmt.Println(pi)

	
}