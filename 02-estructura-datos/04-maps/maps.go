package main

import "fmt"

func main(){
	dias := make(map[int]string)

	fmt.Println(dias)

	// Agregar datos a la estructura de datos de tipo Map dias
	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miercoles"
	dias[4] = "Jueves"
	dias[5] = "Viernes"
	dias[6] = "Sabado"
	dias[7] = "Domingo"

	fmt.Println(dias)

	// Editar un Elemento del Map
	dias[6] = "SABADO"
	fmt.Println(dias)

	// Eliminar un elemento del Map
	delete(dias, 1)
	// No podemos ver la capacidad ya que un mapa no tiene esta propiedad
	fmt.Println(dias, len(dias))

	// Un mapa puede ser una estructura compleja de datos, donde el valor de un indice del mapa
	// puede ser un arreglo, interface, o cualquier tipo de estructura de datos.
	// En este ejemplo, vamos a crear la variable estudiantes mediante la funcion make()
	// y le decimos que el indice va a ser de tipo string, y que el valor va a ser un arreglo de int
	estudiantes := make(map[string][]int)

	estudiantes["Raul"]= []int{13,15,16}
	estudiantes["Alejandro"]= []int{14,18,17}

	fmt.Println(estudiantes)

	fmt.Println(estudiantes["Raul"])

	fmt.Println(estudiantes["Raul"][1])
}