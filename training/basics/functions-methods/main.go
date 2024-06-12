package main

import (
	"fmt"
	"golang-test/training/basics/functions-methods/functions"
	"golang-test/training/basics/functions-methods/methods"
)

func main() {
	//Los métodos son funciones que están asociadas con un tipo de dato específico (estructura).
	//Son definidos dentro de la definición de la estructura.
	//Pueden acceder y modificar los campos de la estructura a la que pertenecen.
	//Se invocan en instancias de la estructura a través del operador .
	rect := methods.Rectangle{Width: 5, Height: 3}
	area := rect.Area()
	fmt.Println("The area of the rectangle is:", area)

	//Las funciones en Go son bloques de código independientes que realizan una tarea específica.
	//Pueden ser definidas fuera de cualquier tipo de estructura.
	//Pueden tomar parámetros y devolver valores.
	//Pueden ser llamadas desde cualquier lugar dentro del paquete en el que están definidas.
	result := functions.Sum(3, 4)
	fmt.Println("El resultado de la suma es:", result)
}
