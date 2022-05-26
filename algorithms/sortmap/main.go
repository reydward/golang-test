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

func main()  {
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

	names := make([]string, 0, len(scores))
	for name := range scores {
		names = append(names, name)
	}

	sort.Slice(names, func(i, j int) bool {
		return scores[names[i]] > scores[names[j]]
	})

	for _, name := range names {
		fmt.Printf("%-7v %v\n", name, scores[name])
	}

//	sortByKey(m)
}

func sortByKey(m map[string]int) map[string]int{
	var sortedMap map[string]int
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}

	return sortedMap
}
