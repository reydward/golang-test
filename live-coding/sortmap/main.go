package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func main() {
	scores := map[string]int{
		"Australia":  24982688,
		"Qatar":      2781677,
		"Wales":      3139000,
		"Burundi":    11175378,
		"Guinea":     12414318,
		"Niger":      22442948,
		"Brazil":     209469333,
		"Malta":      484630,
		"Peru":       31989256,
		"Yemen":      28498687,
		"Ireland":    4867309,
		"Kenya":      51393010,
		"Montserrat": 5900,
		"Cuba":       11338138,
		"Nicaragua":  6465513,
		"Jordan":     9956011,
		"Gabon":      2119275,
	}

	// 1. Crear un slice auxiliar
	type countryScore struct {
		name  string
		score int
	}
	var sortedScores []countryScore

	// 2. Llenar el slice
	for name, score := range scores {
		sortedScores = append(sortedScores, countryScore{name, score})
	}

	// 3. Ordenar el slice
	sort.Slice(sortedScores, func(i, j int) bool {
		return sortedScores[i].name < sortedScores[j].name
	})

	// 4. Imprimir los resultados ordenados
	fmt.Println("Países ordenados por nombre:")
	for _, cs := range sortedScores {
		fmt.Printf("%s: %d\n", cs.name, cs.score)
	}

	//	sortedMap := sortByKey(scores)
	//sortByKey(scores)

	/*	for key, value := range sortedMap {
		println(key, value)
	}*/
}

func sortByKey(m map[string]int) map[string]int {
	sortedMap := make(map[string]int)
	keys := make([]string, 0, len(m))

	// sort the keys
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
		//sortedMap[k] = m[k]
	}

	return sortedMap
}
