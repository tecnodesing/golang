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
	n1, n2, n3 := 8.5, 9.0, 7.5

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


	// Declaración de variables
	helloMessage := "!Hola"
	worldMessage := "Mundo!"

	// Println: Hace un salto de linea al terminar el mensaje
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)
	fmt.Println("")

	// Printf: Imprime con un formato
	nombre := "Platzi"
	cursos := 700
	// Cuando sabemos que tipo de variables son podemos usar
	/*
		 bool:                    %t
		 int, int8 etc.:          %d
		 uint, uint8 etc.:        %d, %#x if printed with %#v
		 float32, complex64, etc: %g
		 string:                  %s
		 chan:                    %p
		 pointer:                 %p
	        */
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)

	// Si desconocemos el tipo de variable en el formato
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf: Genera un string pero no lo imprime en consola
	var message string = fmt.Sprintf("%s es un string", nombre)
	fmt.Println(message)
	fmt.Println("")

	// Imprimir el tipo de variable
	fmt.Printf("HelloMessage es tipo: %T\n", helloMessage)
	fmt.Printf("Cursos es tipo: %T\n", cursos)
	fmt.Println("")

	// Si queremos ignorar el % del formato podemos usar doble %%
	fmt.Printf("%s es de tipo %T y para usarlo en el Printf se hace uso del %%s\n", nombre, nombre)

	



func areaCirculo(radio float64) float64{ 
	return math.Pi*radio*radio
}
func areaRectangulo(base float64, altura float64) float64 {
	return base*altura
}

func areaTrapezoide (B float64,b float64,h float64) float64{
	return h*(B+b)/2
}

func main() {
	fmt.Printf("Circulo %.2f \n",areaCirculo(2))
	fmt.Printf("Rectangulo %.2f \n",areaRectangulo(5,10))
	fmt.Printf("Trapezoide %.2f \n",areaTrapezoide(10,5,3))

}

}
