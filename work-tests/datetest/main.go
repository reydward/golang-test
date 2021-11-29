package main

import (
	"fmt"
	"time"
)

func main()  {

	var bundlingWindowDuration int = 7200
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)
	twoHours := now.Add(time.Hour * -2)
	twoHoursMins := now.Add(time.Second * -7200)
	bundlingWindowLimitFrom := now.Add(time.Second * -time.Duration(bundlingWindowDuration + 600))
	bundlingWindowLimitTo := now.Add(time.Second * -time.Duration(bundlingWindowDuration))

	fmt.Println(yesterday)
	fmt.Println(tomorrow)
	fmt.Println(twoHours)
	fmt.Println(twoHoursMins)
	fmt.Println(bundlingWindowLimitFrom)
	fmt.Println(bundlingWindowLimitTo)

	//layout := "1900-01-01T12:01:28Z"
	updatedAt := "2021-07-26T16:00:28Z"
	updateDate, err := time.Parse(time.RFC3339, updatedAt)
	if err != nil {
		fmt.Println(err)
	}
	diff := now.Sub(updateDate)
	mins := int(diff.Minutes())
	fmt.Printf("Difference in Mins between now: %s and updated: %s, %d Mins\n", now, updateDate, mins)
}