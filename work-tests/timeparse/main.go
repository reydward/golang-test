package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func isExpired(emailUserCreatedAt string) (bool, error) {
	today := time.Now()
	createdAt, err := time.Parse(time.RFC3339, emailUserCreatedAt)

	if err != nil {
		return true, err
	}

	cutOffDay := createdAt.Add(24 * time.Hour)

	if today.After(cutOffDay) {
		return true, errors.New("token has expired")
	}

	return false, nil
}
func main() {
	//	currentTime := time.Now()

	//authExpired, err := isExpired("2023-04-28T15:37:42.514862Z")
	//createdAt, err := time.Parse(time.RFC3339, "2023-04-28T15:37:42.514862Z")
	/*
		valuationDate := time.Now().Format(time.RFC3339)[0:10]
		valuationDateFormatted := valuationDate
		//time.DateOnly
		fmt.Println(valuationDate)
		fmt.Println(valuationDateFormatted)
	*/
	ValuationDate := "2023/04/01"
	res := strings.Split(ValuationDate, "/")
	if len(res) == 3 {
		fmt.Println(res[0] + "-" + res[1] + "-" + res[2])
	} else {
		fmt.Println("Wrong format")

	}
}
