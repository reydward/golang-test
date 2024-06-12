package main

func main() {
/*	type1 := type1{
		"type1 value",
	}
	printMessage(type1)
	type2 := type2{
		"type2 value",
	}
	printMessage(type2)
*/
/*
	musicalplayers := []MusicalPlayer {
		Trumpeter{
			"trumpeter 1",
		},
		Trumpeter{
			"trumpeter 2",
		},
		Violinist{
			"violinist 1",
		},
	}

	for _, x := range musicalplayers{
		x.Greetings()
	}
*/
/*
	circle := Circle{
		777,
	}
	ExecuteArea(circle)

	rectangle := Rectangle{
		777,
		333,
	}
	ExecuteArea(rectangle)

 */

	book1 := Book{
		"book 1",
		100,
	}
	game1 := Game{
		"game 1",
		101,
	}

	ExecuteProduct(&book1, 10)

	ExecuteProduct(&game1, 20)
}


