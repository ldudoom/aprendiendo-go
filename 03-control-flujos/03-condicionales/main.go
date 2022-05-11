package main

import "fmt"

func main(){
	/** Ap para restaurante
	*	Descuentos de 10% hasta 100 de consumo
	*	Descuentos de 20% hasta 200 de consumo
	*	Descuentos de 30% mayor que 200 de consumo
	*	iva 12%
	**/

	var consumo, descuento float64
	var datosDescuento string

	// Entrada de Datos
	fmt.Println("Ingrese el consumo: ")
	fmt.Scanln(&consumo)

	if consumo > 0 {
		if consumo > 0 && consumo <= 100 {
			datosDescuento = "10%"
			descuento = consumo * (10.0 / 100.0) // Se colocan valores decimales ya que la variable descuento es de tipo float64
		} else if consumo > 100 && consumo <= 200 {
			datosDescuento = "20%"
			descuento = consumo * (20.0 / 100.0) // Se colocan valores decimales ya que la variable descuento es de tipo float64
		} else if consumo > 200 {
			datosDescuento = "30%"
			descuento = consumo * (30.0 / 100.0) // Se colocan valores decimales ya que la variable descuento es de tipo float64
		}

		subtotal := consumo - descuento

		iva := subtotal * (12.0/100.0) // Se colocan valores decimales ya que la variable iva es de tipo float64

		total := subtotal + iva

		fmt.Printf("Su descuento es del %s \n", datosDescuento)
		fmt.Printf("Precio: %v \n", consumo)
		fmt.Printf("Descuento: %v \n", descuento)
		fmt.Printf("Subtotal: %v \n", subtotal)
		fmt.Printf("IVA: %f \n", iva)
		fmt.Printf("Total: %v \n", total)
	}else{
		datosDescuento = "Por favor ingrese un valor de consumo válido"
		fmt.Println(datosDescuento)
	}
	


}