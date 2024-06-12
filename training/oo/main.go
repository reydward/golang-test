package main

import "fmt"

// Struct para representar un producto
type Producto struct {
	Nombre string
	Precio float64
}

// Métodos para el struct Producto
func (p *Producto) Mostrar() {
	fmt.Println("Nombre:", p.Nombre)
	fmt.Println("Precio:", p.Precio)
}

func (p *Producto) CalcularIVA() float64 {
	return p.Precio * 0.21
}

// Struct para representar un libro (hereda de Producto)
type Libro struct {
	Producto // Herencia embedida
	Autor    string
}

// Método único para Libro
func (l *Libro) MostrarAutor() {
	fmt.Println("Autor:", l.Autor)
}

func main() {
	// Crear un producto
	producto := Producto{"Laptop", 1200.00}
	producto.Mostrar() // Invoca al método Mostrar() del struct Producto

	// Crear un libro y heredar atributos y métodos de Producto
	libro := Libro{Producto{"Libro de Go", 35.99}, "Alan Turing"}
	libro.Mostrar()                          // Invoca al método Mostrar() heredado de Producto
	libro.MostrarAutor()                     // Invoca al método único de Libro
	fmt.Println("IVA:", libro.CalcularIVA()) // Invoca al método heredado de Producto
}
