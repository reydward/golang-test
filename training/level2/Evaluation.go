package main

import "fmt"

/*
A store sells two types of products: books and games. Each product has the fields name (string) and price (float). Define the following functionalities:
Product has the following methods:
A method to print the information of each product (type, name, and price)
A method to apply a discount ratio to the price
The store should be able to apply custom discounts based on the type of product: 10% discount for books and 20% discount for games.
 */

type Product interface {
	GetInformation()
	ApplyDiscount(percentage float64)
}

type Book struct {
	name string
	price float64
}

type Game struct {
	name string
	price float64
}

func (b Book) GetInformation() {
	fmt.Println("Book", b.name, b.price)
}

func (b *Book) ApplyDiscount(percentage float64) {
	b.price = b.price - (b.price * (percentage/100))
}

func (g Game) GetInformation() {
	fmt.Println("Game", g.name, g.price)
}

func (g *Game) ApplyDiscount(percentage float64) {
	g.price = g.price - (g.price * (percentage/100))
}

func ExecuteProduct(p Product, percentage float64) {
	p.GetInformation()
	p.ApplyDiscount(percentage)
	p.GetInformation()
}

