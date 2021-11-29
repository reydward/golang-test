package main

import "fmt"

var (
	m = make(map[string]int)
)

func main() {
	m["Creamy Dill Chicken"] = 1
	m["Speedy Steak Fajitas"] = 2

	s := []string{"Creamy Dill Chicken", "Creamy Dill Chicken", "Speedy Steak Fajitas", "test", "test"}

	for _, recipe := range s {
		uniqueRecipeCount(recipe)
	}

	fmt.Println(m)
}

func uniqueRecipeCount(recipe string) {
	_, exists := m[recipe]
	if exists {
		m[recipe] = m[recipe] + 1
	} else {
		m[recipe] = 1
	}
}
