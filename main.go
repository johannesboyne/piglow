package main

import (
	"os"

	"github.com/johannesboyne/gopiglowinterface"
)

func main() {
	color := os.Args[1:][0]
	glower := gopiglowinterface.Glow{}
	glower.NewGlow()

	switch color {
	case "green":
		glower.Green(100)
	case "red":
		glower.Red(100)
	case "orange":
		glower.Orange(100)
	case "blue":
		glower.Blue(100)
	case "clear":
		glower.Clear()
	}
}
