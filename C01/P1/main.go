// Calcular la media de 2 numeros
package main

import (
	"fmt"
)

func main() {
	var num1, num2 int

	num1 = 10
	num2 = 15

	result := (num1 + num2) / 2

	fmt.Println("La media de", num1, "y", num2, "es", result)
}