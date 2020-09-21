package main

import "fmt"

func main() {

	var f float64

	fmt.Printf("Convertir grados fahrenheit a celcius\n\n")
	fmt.Print("Ingresa grados fahrenheit: ")
	fmt.Scanf("%f", &f)
	celcius := (f - 32) * 5 / 9
	fmt.Println("\nGrados celcius: ", celcius)

}
