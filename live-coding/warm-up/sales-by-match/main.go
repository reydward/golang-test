package main

import (
	"fmt"
)

/*
There is a large pile of socks that must be paired by color. Given an array of integers representing the color of each sock,
determine how many pairs of socks with matching colors there are.
Example: There is one pair of color and one of color.
There are three odd socks left, one of each color. The number of pairs is .

 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
*/

func sockMerchant(n int32, socks []int32) int32 {
	// Creamos un mapa para contar la cantidad de cada color de calcetín
	counts := make(map[int32]int32)

	// Contamos la cantidad de calcetines de cada color
	for _, color := range socks {
		counts[color]++
	}

	// Inicializamos el contador de pares
	pairs := int32(0)

	// Iteramos sobre los conteos de calcetines de cada color
	for _, count := range counts {
		// Calculamos cuántos pares hay para cada color de calcetín
		pairs += count / 2
	}

	return pairs
}

func main() {
	// Ejemplo de uso
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	n := int32(len(ar))
	result := sockMerchant(n, ar)
	fmt.Println("Número de pares de calcetines:", result)
}
