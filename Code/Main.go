package main

import (
	"fmt"
	"strings"
)

func main() {


	//El metodo Println sirve para imprimir un texto en consola
	fmt.Println("Hello World")

	// messagesfromdoris := []string{
	// 	"Hola Mundo",
	// 	"Adios Mundo",
	// 	"EL Precio de esto es de",
	// }

	// // numMessages := float64(len(messagesfromdoris))
	// // costMessage := 3.0
	// // Total := numMessages * costMessage
	// // fmt.Println(Total)


	// var username string = "Misael"
	// var password int= 1234567
	// var pi float32= 3.1416
	// var condtion bool = true
	// var age uint8= 18

	// fmt.Println()

// edad := 15

// if edad <= 17 {
// 	fmt.Println("Es un niño")
// } else if edad >= 18 {
// 	fmt.Println("Es un adolecente")
// }


// var numbers [3] int= [3]int{1,2,3}

// fmt.Println(numbers[2])
	
//Arrays
// numbers:= [4]string {"Hola","Mundo","Como","Estas?"}

// fmt.Println(numbers[0] + " " + numbers[1] + " " + numbers[2] + " " + numbers[3])

//Slices o Arrays Dinamicos
nombre:= "Misael Alejandro"
var numbers= []int{1,2,3}
numbers[2]= 4;
numbers= append(numbers, 5)

fmt.Println(numbers,len(numbers))

// ranges 

range_array:= numbers[0:2]
range_init:= numbers[2:]
range_end:=numbers[:4]

fmt.Println(range_array)
fmt.Println(range_init)
fmt.Println(range_end)
// El package strings es para trabajar con cadenas de texto
fmt.Println(strings.Contains(nombre,"Misael"))
//Remplaza una parte de la Cadena de Texto
fmt.Println(strings.ReplaceAll(nombre,"Misael","Misa"))
// Escribir todo en Mayuscula
fmt.Println(strings.ToUpper(nombre))
// Encontrar el caracter almacenado en cierta posicion
fmt.Println(strings.Index(nombre,"Ale"))

}
