package main

import (
	"fmt"
	"os"
)

func main() {

	/*
		La funcion recover lo que hace es controlar todos los panic, de tal manera que si ocurre
		algun error en la aplicacion, lo podamos tener controlado y la aplicacion no se cerrara de
		manera brusca. Y la vamos a ejecutar junto con un defer para que se ejecute al final
	*/

	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Ups! al parecer el programa no finalizo de forma correcta: ", error)
		}
	}()

	// Vamos a leer el archivo hola.txt
	// vamos a declarar 2 variables, una para obtener la informacion del archivo y otra para obtener el error
	// aqui vamos a utilizar la libreria OS que sirve para realizar acciones con metodos del sistema
	//file, error := os.Open("./hola.txt")
	//if error == nil{
	// Cambiamos las 2 instrucciones anteriores por el siguiente if
	if file, error := os.Open("./holas.txt"); error == nil {
		// Ahora para cerrar el archivo debemos utilar defer, que lo que hace es hacer que el siguiente bloque de
		// codigo se ejecute al final de las instrucciones
		defer func() {
			fmt.Println("Cerrando el archivo")
			file.Close()
		}()

		// Creamos un slice vacio para obtener la informacion del archivo aqui dentro
		contenido := make([]byte, 254)
		// declaramos una variable en la cual se va a almacenar la longitud del archivo,
		// y como segundo parametro vamos a colocar un "_" ya que no queremos manejar el error este momento
		// por ultimo, utilizamos la funcion Read() de file y le pasamos como parametro nuestro slice
		longitud, _ := file.Read(contenido)
		// En el slice contenido tenemos ya los datos del archivo pero de tipo Byte, entonces lo que vamos a hacer es convertir
		// esos datos para posteriormente poder leerlos. Para eso, declaramos una nueva variable, y dentro de esta almacenamos
		// el resultado de la conversion.
		// Para eso, utilizamos la funcion string, y le pasamos como parametro el slice, diciendole que lo lea desde el inicio
		// hasta el fin que es la longitud del archivo que obtuvimos anteriormente, puede hacerse de estas 2 maneras:
		// contenido[:longitud]   /    contenido[0:longitud]    estas 2 maneras indican exactamente la misma accion
		contenidoArchivo := string(contenido[:longitud])
		fmt.Println(contenidoArchivo)
	} else {
		//fmt.Println("No pudimos leer el archivo debido a:", error)
		// La funcion panic detiene el proceso cuando encuentra un error y dispara un error junto
		// con el mensaje que le indicamos
		result := fmt.Sprintln("No pudimos leer el archivo debido a: ", error)
		panic(result)
	}

	// Por ultimo, cerramos el archivo con el que estuvimos trabajando
	//file.Close()

}
