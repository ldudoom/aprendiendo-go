package main

import "fmt"

func main(){
	// Slice: son parecidos a los arreglos, la diferencia es que estos si son mutables
	numeros := []int{1,2,3}
	fmt.Println(numeros, len(numeros))

	// Agregar elementos
	numeros = append(numeros, 4)
	numeros = append(numeros, 5)
	fmt.Println(numeros, len(numeros))

	// Sub Slice
	subSlice := numeros[:2]
	numeros[0] = 100
	fmt.Println(subSlice, len(subSlice))
	fmt.Println(numeros, len(numeros))

	// Para darse cuenta que una estructura de datos es un Slice, éste debe tener 3 características
	// Tener un puntero a un array, 
	// Una longitud dinámica
	// Tener una Capacidad

	// Punteros -> subSlicen := numeros[:2]
	// Longitud -> len(numeros)
	// Capacidad -> cap(numeros)

	/*
		Una de las más importantes características de los slices es que permiten aumentar la capacidad 
		máxima de elementos por medio de la función append() la cual añade uno o varios elementos nuevos 
		desde el último índice del slice. Si el slice aún tiene sus últimos índices disponibles, 
		simplemente se añade el nuevo elemento a la cola; en cambio, si el último índice ya está utilizado, 
		se aumenta dinámicamente la capacidad del slice al doble y se añade dicho elemento en la última 
		posición.
	*/

	meses := []string{"Enero", "Febrero", "Marzo"}
	fmt.Printf("Longitud: %v, Capacidad: %v, %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "Abril")
	fmt.Printf("Longitud: %v, Capacidad: %v, %p \n", len(meses), cap(meses), meses)

	/*
		El valor cero de un slice es nil
		Un slice nil tiene una longitud y capacidad de 0 y no tiene un arreglo subyacente
	*/

	var slice []int
	fmt.Println(slice, len(slice), cap(slice))
	if slice == nil {
		fmt.Println("nil!")
	}



}