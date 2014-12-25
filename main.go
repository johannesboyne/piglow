package main

import (
	"fmt"
	"os"

	"github.com/johannesboyne/gopiglowinterface"
)

func main() {
	color := os.Args[1:][0]
	glower := gopiglowinterface.Glow{}
	glower.NewGlow()

	switch color {
	case "green":
		fmt.Println("GRÜN")
		glower.Green(100)
	case "red":
		fmt.Println("ROT")
		glower.Red(100)
	case "orange":
		fmt.Println("ORANGE")
		glower.Orange(100)
	case "blue":
		fmt.Println("BLAU")
		glower.Blue(100)
	}
}
