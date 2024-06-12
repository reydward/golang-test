package structures

type Product interface {
	GetID() int
	GetName() string
	GetPrice() float64
	GetQuantity() int
	ApplyDiscount(percentage float64)
}

type Book struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

type Game struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}
