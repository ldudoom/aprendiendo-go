package main

import "fmt"

func main(){
	hola := "Hola"
	mundo := "Mundo"

	// el metodo Print imprime en pantalla sin saltos de linea
	fmt.Print(hola)
	fmt.Print(mundo)

	// EL metodo Println imrpime en pantalla y agrega saltos de linea
	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Raul"
	edad := 40
	// El metodo Printf sitve para formatear el valor de las variables
	// Si nosotros vamos a imprimir una variable que tiene un string colocamos %s, si va a ser entero %d
	fmt.Printf("Hola %s, tu edad es %d años\n", nombre, edad)
	
	// Si nosotros no sabemos que tipo de dato tiene la variable, y por lo tanto no sabemos que se va a imprimir
	// entonces utilizamos %v
	fmt.Printf("Hola %v, tu edad es %v años\n", nombre, edad)

	// EL metodo Sprintf nos permite almacenar en una variable la cadena con las variables formateadas
	mensaje := fmt.Sprintf("Hola %s, tu edad es %d años\n", nombre, edad)
	fmt.Println(mensaje)

	// Para conocer que tipo de dato contiene la variable nombre podemos utilizar Printf de la siguiente manera
	fmt.Printf("nombre: %T \n edad: %T \n", nombre, edad)

	// Vamos a imprimir el mensaje para que el usuario ingrese su apellido y lo vamos a almacenar en una variabñe
	var apellido string
	fmt.Print("Ingrese su apellido: ")
	fmt.Scanln(&apellido)

	fmt.Println("Apellido: ", apellido)


}