package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("¡Hola, Mundo!")
	color.Cyan("¡Hola %s", "Go")
}
