package main

import "fmt"

func main(){
	nombres := [...]string{"Raul","Alejandro","Jacqueline","Fernanda","Isaac","Nicole","Anahi","Astrid"}

	// Usando un for normal
	fmt.Println("Con un for normal")
	for i:=0;i<len(nombres);i++ {
		fmt.Println(nombres[i])
	}
	fmt.Println("*************************************************************")

	// Usando for para hacer un foreach
	fmt.Println("Haciendo funcionar un for como si fuera un foreach")
	for key, value := range nombres {
		fmt.Println(key, value)
	}

	fmt.Println("*************************************************************")
	// Usando for para hacer un foreach
	fmt.Println("Si no queremos imprimir el indice")
	for _, value := range nombres {
		fmt.Println(value)
	}
}