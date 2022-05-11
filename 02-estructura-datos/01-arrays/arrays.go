package main

import "fmt"

func main(){
	var numeros [5]int // definimos la longitud del arreglo la cual es inmutable, es decir no se podra cambiar
	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	fmt.Println(numeros)
	fmt.Println(numeros[1])

	// Definir un arreglo y asignarle los valores que va a tener cada indice del arreglo
	nombres := [3]string{"Raul", "Fher", "Astrid"} // definimos la longitud en 3
	fmt.Println(nombres)

	// se puede definir un arreglo sin definir la longitud, y colocar los valores por defecto que requiramos
	// Sin embargo, una vez que el arreglo esta definido, y sus valores asignados, esta longitud no puede cambiar 
	colores := [...]string{
		"Azul",
		"Negro",
		"Rojo",
		"Verde",
		"Amarillo",
		"Violeta",
		"Blanco",
		"Indigo",
		"Naranja"}
	fmt.Println(colores, len(colores))

	// Vamos a definir los indices
	monedas := [...]string{
		0:"Dolar",
		2:"Euros",
		3:"Pesos",
		5:"Soles",
	}
	fmt.Println(monedas, len(monedas))

	// Sub array
	subMonedas := monedas[0:3]
	// la linea de arriba es igual a decirle:
	// subMonedas := monedas[:3]
	fmt.Println(subMonedas, len(subMonedas))


}