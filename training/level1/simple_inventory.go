/*
Simple Inventory
Statement
Create a Product struct with an attribute ID and Name, and an Inventory struct that will contain a map[string]Products.
Create a method with a pointer of Inventory as a receiver and It will expect a product to be added into the map of products.
Required: Add validations into the Add method, to check that the ID can’t be empty and also not duplicated ID

Topics to Practice:
struct, composition, methods, pointer, error handling
*/

package level1

type Product struct {
	ID   int
	Name string
}

type Inventory struct {
	Products map[string]Product
}

func (inventory *Inventory) AddProduct(product Product) {

}
