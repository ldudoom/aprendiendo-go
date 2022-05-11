package main

import "fmt"

func main(){
	
	verify := true
	// Not
	fmt.Println( ! verify)

	// And
	fmt.Println(verify && verify)
	fmt.Println(verify && !verify)
	fmt.Println(!verify && verify)
	fmt.Println(!verify && !verify)

	// Or
	fmt.Println(verify || verify)
	fmt.Println(verify || !verify)
	fmt.Println(!verify || verify)
	fmt.Println(!verify || !verify)

	b := 2
	r := b == 2 && 4 > b
	fmt.Println(r)

	r = b == 2 && !(4 > b)
	fmt.Println(r)

	r = b == 2 || !(4 > b)
	fmt.Println(r)
}