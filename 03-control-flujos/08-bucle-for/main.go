package main

import "fmt"

func main(){
	// Bucle infinito
	/*
	for{
		fmt.Println("Te van a contratar en Canada :) !!!!")
	}*/	

	// Bucle tipo while
	numeros := 12455123
	cont := 0
	for numeros > 0 {
		numeros /= 10
		cont++
	}
	fmt.Println("Cantidad de digitos es:", cont)

	// Bucle for
	for i := 0; i <= 100; i++{
		if i%2 == 0{
			fmt.Println(i)
		}
	}

}