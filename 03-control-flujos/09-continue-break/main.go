package main

import "fmt"

func main(){
	for i:=0;i<=10;i++ {
		if i == 8 {
			fmt.Println("Rompemos el bucle en i =", i)
			break;
		}
		if i == 7 {
			fmt.Println("Con i =", i,"continuamos el bucle saltandonos las acciones con este valor")
			continue;
		}
		fmt.Println(i)
	}
}