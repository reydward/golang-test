package main

/**
In a card game, each player will be given a set of random cards. Players will throw on the table their one winning card, the player with the highest card wins.
A winning card is the card that only exists once in the set of cards, and the highest one.

Given an array of sets of integers cards, return the card of the winning player. Return -1 If no such card is found.

Example 1:
Input: cards = [[5,7,3,9,4,9,8,3,1], [1,2,2,4,4,1], [1,2,3]]
Output: 8

Example 2:
Input: cards = [[5,5], [2,2]]
Output: -1

Constraints:
1 <= cards.length <= 2000
0 <= cards[i] <= 1000
*/
import "fmt"

func winningCard(cards [][]int) int {
	// Creamos un mapa para contar la frecuencia de cada carta
	freq := make(map[int]int)

	// Iteramos sobre cada conjunto de cartas y actualizamos la frecuencia de cada carta
	for _, hand := range cards {
		for _, card := range hand {
			freq[card]++
		}
	}

	// Inicializamos el valor de la carta ganadora
	//maxFreq := 0
	winningCard := -1

	// Iteramos sobre las frecuencias y encontramos la carta ganadora
	for card, count := range freq {
		println("Frecuencia de cartas:", card, count)
		if count == 1 && card > winningCard {
			winningCard = card
			fmt.Println("Carta ganadora por el momento:", winningCard)
		}
	}

	return winningCard
}

func main() {
	// Ejemplo de uso
	cards1 := [][]int{{5, 7, 3, 9, 4, 9, 8, 3, 1}, {1, 2, 2, 4, 4, 1}, {1, 2, 3}}
	fmt.Println("Carta ganadora (Ejemplo 1):", winningCard(cards1))

	cards2 := [][]int{{5, 5}, {2, 2}}
	fmt.Println("Carta ganadora (Ejemplo 2):", winningCard(cards2))

	cards3 := [][]int{{5, 5, 7}, {2, 2, 9}}
	fmt.Println("Carta ganadora (Ejemplo 3):", winningCard(cards3))
}
