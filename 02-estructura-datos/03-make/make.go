package main

import (
	"fmt"
	"strings"
)

/*
	The make built-in function allocates and initializes an object of type
	slice, map, or chan (only). Like new, the first argument is a type, not a
	value. Unlike new, make's return type is the same as the type of its
	argument, not a pointer to it. The specification of the result depends on
	the type:
		Slice: The size specifies the length. The capacity of the slice is
		equal to its length. A second integer argument may be provided to
		specify a different capacity; it must be no smaller than the
		length. For example, make([]int, 0, 10) allocates an underlying array
		of size 10 and returns a slice of length 0 and capacity 10 that is
		backed by this underlying array.
		Map: An empty map is allocated with enough space to hold the
		specified number of elements. The size may be omitted, in which case
		a small starting size is allocated.
		Channel: The channel's buffer is initialized with the specified
		buffer capacity. If zero, or the size is omitted, the channel is
		unbuffered.
*/

func main(){
	//numeros := make([]int, 0,3)
	/* 
	 El siguiente bloque de código da error ya que lo que estoy haciendo es modificar los indices, no
	 agregarles valores
	numeros[0] = 100
	numeros[1] = 200
	numeros[2] = 300

	Para que esto funcione, lo que debemos hacer es cambiar esto
		numeros := make([]int, 0,3)
	por esto:
		numeros := make([]int, 3,3)
	*/

	numeros := make([]int, 3,3)
	
	numeros[0] = 100
	numeros[1] = 200
	numeros[2] = 300
	
	fmt.Println(numeros, len(numeros), cap(numeros))

	numeros = append(numeros, 400)

	fmt.Println(numeros, len(numeros), cap(numeros))

	numeros = append(numeros, 500)

	fmt.Println(numeros, len(numeros), cap(numeros))

	numeros = append(numeros, 600)
	numeros = append(numeros, 700)
	numeros = append(numeros, 800)
	numeros = append(numeros, 900)

	fmt.Println(numeros, len(numeros), cap(numeros))

	subSlice := numeros[2:]
	fmt.Println(subSlice, len(subSlice))

	numeros = append(numeros, 900)

	// Slice de Slices
	// Los Slices pueden contener elementos de cualquier tipo, incluyendo otros slices.

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	/*
		[]string{"X", "_", "X"},
		[]string{"O", "_", "X"},
		[]string{"_", "_", "O"},
	*/

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}


	// RANGE
	/*
		la forma range en un bucle for se itera sobre un slice o un map
		Cuando rage se itera sobre un elemento, se devuelven 2 valores en cada iteración.
		El primero es el índice del slice o del map, y el segundo es una copia del elemento de ese indice
	*/

	var potencia = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// En este caso, cuando obtenermos el range del slice potencia dentro del bucle for
	// en la variable "i" se almacena el indice del slice, y en la variable "v" se guarda una 
	// copia del valor del indice "i"
	for i, v := range potencia {
		fmt.Printf("2^%d = %d\n", i, v)
	}

}