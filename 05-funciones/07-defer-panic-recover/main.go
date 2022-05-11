package main

import (
	"fmt"
	"os"
)

func main(){
	// Vamos a leer el archivo hola.txt
	// vamos a declarar 2 variables, una para obtener la informacion del archivo y otra para obtener el error
	// aqui vamos a utilizar la libreria OS que sirve para realizar acciones con metodos del sistema
	file, error := os.Open("./hola.txt")
	if error == nil{
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
		fmt.Println("No pudimos leer el archivo debido a:", error)
	}

	// Por ultimo, cerramos el archivo con el que estuvimos trabajando
	file.Close()
	
}