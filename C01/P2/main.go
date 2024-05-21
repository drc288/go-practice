// Condicionales
package main

import (
	"fmt"
)

func main() {
	var num1, num2 int

	num1 = 15
	num2 = 20

	if num1 > num2 {
		fmt.Println(num1, "es mayor que", num2)
	} else {
		fmt.Println(num1, "es menor que", num2)
	}
}