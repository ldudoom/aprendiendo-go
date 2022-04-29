package main

import (
	"errors"
	"fmt"
)

// Esta funcion va a devolver 2 valores, el uno va a ser de tipo entero, y el otro de un tipo de dato error

func division(dividendo, divisor int) (int, error){
	if divisor == 0{
		// En caso de que se envie el numero 0 como divisor, controlamos la excepcion manejando nosotros el error
		// retornamos el valor entero 0 y un error con un mensaje personalizado
		return 0, errors.New("No es posible dividir un numero entre 0")
	} else {
		// En caso de que se envien valores correctos, devolvemos el resultado de la division, y un nil como segundo valor
		// ya que no existio ningun error
		return dividendo / divisor, nil
	}
}

func main(){
	// recibimos el resultado de la funcion dividir en las variables result y error
	result, error := division(4, 0)
	// En caso de que error tenga el valor de nil, imprimimos el resultado de la funcion dividir
	// Caso contrario, imprimimos el error
	if error == nil{
		fmt.Println(result)
	}else{
		fmt.Println(error)
	}
	
}