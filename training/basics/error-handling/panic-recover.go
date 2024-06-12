package main

import "fmt"

func main() {
	//recuperar el control después de que un programa ha entrado en un estado de pánico
	//Se utiliza dentro de una función diferida (defer)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado:", r)
		}
	}()

	fmt.Println("Iniciando el pánico...")
	panic("¡Oh no! Ha ocurrido un error crítico.")
	fmt.Println("Este mensaje nunca se imprimirá debido al pánico.")
}
