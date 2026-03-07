package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("¡Hola, Mundo!")
	color.Cyan("¡Hola %s", "Go")

	var mensaje string
	mensaje = "Hola mundo"
	var edad int
	var mensaje2 = " años"
	edad = 31
	mensaje += "Mi edad es: "
	fmt.Println(mensaje, edad, mensaje2)
	var mensaje3 string
	mensaje3 = fmt.Sprintf("%s%d%s", mensaje, edad, mensaje2)
	fmt.Println(mensaje3)
}
