package main

import (
	"sync"
)

/*
Cálculo Concurrente de Sumas Parciales:
Dado un slice de enteros, divide el slice en partes iguales y calcula la suma de cada parte concurrentemente. Finalmente, combina las sumas parciales para obtener la suma total del slice.

Requisitos:

Divide el slice en n partes iguales (o casi iguales si la longitud del slice no es divisible por n).
Calcula la suma de cada parte concurrentemente usando goroutines.
Usa un canal para recoger las sumas parciales.
Imprime la suma total del slice.
*/

// Función para calcular la suma de una parte del slice
func sumPart(nums []int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, num := range nums {
		sum += num
	}
	result <- sum
}

func main() {
	// Slice de enteros a sumar
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Número de partes en las que se dividirá el slice
	n := 3

	// Canal para recoger los resultados de las sumas parciales
	result := make(chan int, n)
	var wg sync.WaitGroup

	// Determinar el tamaño de cada parte
	partSize := (len(nums) + n - 1) / n

	// Dividir el slice y lanzar goroutines
	for i := 0; i < n; i++ {
		start := i * partSize
		end := start + partSize
		if end > len(nums) {
			end = len(nums)
		}
		wg.Add(1)
		go sumPart(nums[start:end], result, &wg)
	}

	// Esperar a que todas las goroutines terminen
	wg.Wait()
	close(result)

	for partialSum := range result {
		println("Channel result:", partialSum)
	}

	// Calcular la suma total
	/*	totalSum := 0
		for partialSum := range result {
			totalSum += partialSum
		}

		// Imprimir la suma total
		fmt.Printf("Total sum of the slice is: %d\n", totalSum)
	*/
}
