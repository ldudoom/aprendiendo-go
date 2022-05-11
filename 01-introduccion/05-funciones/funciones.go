package main

import "fmt"

func main(){
	
	saludo("Hola como estas")
}

func saludo(mensaje string){
	fmt.Println(mensaje, getName())
}

func getName() string{
	var nameUser string
	var lastnameUser string
	fmt.Print("Imgresa tu primer nombre: ")
	fmt.Scanln(&nameUser)
	fmt.Print("Imgresa tu apellido: ")
	fmt.Scanln(&lastnameUser)
	var nombreCompleto string = fmt.Sprintf("%s %s", nameUser, lastnameUser)
	return nombreCompleto
}