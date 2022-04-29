package main 

import(
	"fmt"
	"strings"
	"runtime"
) 

func main (){
	var vocal string

	fmt.Print("Por favor ingrese una letra: ")
	fmt.Scanln(&vocal)

	switch strings.ToLower(vocal){
	case "a":
		fmt.Println(vocal, ": Es vocal")
	case "e":
		fmt.Println(vocal, ": Es vocal")
	case "i":
		fmt.Println(vocal, ": Es vocal")
	case "o":
		fmt.Println(vocal, ": Es vocal")
	case "u":
		fmt.Println(vocal, ": Es vocal")
	default :
		fmt.Println(vocal, ": Es consonante")
	}

	// Se puede simplificar el codigo de esta manera, ya que el texto de la respuesta siempre es el mismo
	switch strings.ToLower(vocal){
	case "a", "e", "i", "o", "u":
		fmt.Println(vocal, ": Es vocal")
	default :
		fmt.Println(vocal, ": Es consonante")
}

	fmt.Println("//////////////////////////////////////////////////////////////")
	fmt.Println(runtime.GOOS)
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	 
}