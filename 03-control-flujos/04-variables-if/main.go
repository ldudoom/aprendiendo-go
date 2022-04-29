package main

import "fmt"

func main(){
	if nombre, edad := "Alejandro", 40; nombre == "Raul" {
		fmt.Println("Hola ", nombre)
	}else{
		fmt.Printf("Nombre: %s - Edad: %d\n", nombre, edad)
	}

	// obtener el valor de un map
	users := make(map[string] string)

	users["Raul"] = "raullduch@gmail.com"
	users["Alejandro"] = "raul.chauvin.ojeda@gmail.com"

	// Cuando intentamos obtener información de un map, lo podemos hacer de la siguiente manera
	correo, ok := users["Raul"]
	// En este caso, el mapa devuelve 2 valores, primero, devuelve el valor de la estructura clave/valor,
	// es decir en este caso el correo, y lo almacenamos en la variable correo, como segundo dato devuelve
	// un valor boolean que indica si existe el elemtno que estamos consultando (true) o si no existe (false)
	fmt.Println(correo, ok)

	// De esta manera podriamos utilizar el mapa en el scope (ambito) del if de la siguiente manera
	// para escribir un mensaje en caso de que el indice exista, y un mensaje de error en caso de
	// que no exista.
	//if correo, ok := users["Juan"]; ok{
	if correo, ok := users["Raul"]; ok{
		fmt.Println(correo)
	}else{
		fmt.Println("Ups! No pudimos obtener el correo.")
	}
}