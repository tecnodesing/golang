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
	fmt.Printf("El valor de pi es: %.2f\n", pi)
}
