package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("¡Hola, Mundo!")
	color.Cyan("¡Hola %s", "Go")

	var mensaje string
	mensaje = "Hola clase de Go. "
	var edad int
	var mensaje2 = " años"
	edad = 31
	mensaje += "Mi edad es: "
	fmt.Println(mensaje, edad, mensaje2)
	var mensaje3 string
	mensaje3 = fmt.Sprintf("%s%d%s", mensaje, edad, mensaje2)
	fmt.Println(mensaje3)

	const mensajeBienvenida = "Bienvenidos al sistema de Go"
	const nombre = "Irving"
	const pi = 3.1416

	fmt.Println(mensajeBienvenida, nombre)
	fmt.Printf("El valor de pi es: %.4f\n", pi)

	nombre2 := "Irving"
	apellido := "velasco"
	n1, n2, n3 := 8.5, 9, 7.5

	//operaciones aritmeticas
	suma := n1 + n2 + n3
	resta := n1 - n2 - n3
	multiplicacion := n1 * n2 * n3
	division := n1 / n2 / n3
	promedio := suma / 3
	nombreCompleto := nombre2 + " " + apellido

	fmt.Printf("El nombre completo es: %s\n", nombreCompleto)
	fmt.Printf("La suma es: %.2f\n", suma)
	fmt.Printf("La resta es: %.2f\n", resta)
	fmt.Printf("La multiplicación es: %.2f\n", multiplicacion)
	fmt.Printf("La división es: %.2f\n", division)
	fmt.Printf("El promedio es: %.2f\n", promedio)

}
