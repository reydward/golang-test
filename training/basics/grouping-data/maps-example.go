package main

import "fmt"

func main() {
	m := map[string]int{
		"Batman": 32,
		"Robin":  27,
	}
	fmt.Println(m)

	fmt.Println(m["Batman"])

	fmt.Println(m["Eduar"])

	v, ok := m["Robin"] //ok es un booleano que indica si la llave existe o no
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Eduar"]; ok {
		fmt.Println("Impresión desde el if", v)
	}

	//Agregar elementos al map
	m["Eduard"] = 333

	for k, v := range m {
		fmt.Println(k, v)
	}

	//Eliminar elementos del map
	delete(m, "Batman")

	for k, v := range m {
		fmt.Println(k, v)
	}

}
