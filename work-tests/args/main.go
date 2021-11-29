package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	inputFile := flag.String("inputFile", "", "Input file")
	postcodeFlag := flag.String("postcode", "", "Number of deliveries to postcodeFlag")
	startWindowFlag := flag.String("startWindow", "", "Start delivery time")
	endWindowFlag := flag.String("endWindow", "", "End delivery time")
	flag.Parse()
	fmt.Println("inputFile:", *inputFile)
	fmt.Println("postcodeFlag:", *postcodeFlag)
	fmt.Println("startWindowFlag:", *startWindowFlag)
	fmt.Println("endWindowFlag:", *endWindowFlag)
	fmt.Println("tail:", flag.Args())

/*	fmt.Println(convertTo24h("11PM"))
	fmt.Println(getStartTime("Wednesday 7AM - 7PM"))
	fmt.Println(getEndTime("Wednesday 7AM - 7PM"))*/

}

func convertTo24h(time string) int {
	intTime := 0
	intTime, _ = strconv.Atoi(time[0:len(time)-2])
	if strings.Contains(time, "PM") {
		intTime = intTime + 12
	}
	return intTime
}


func getEndTime(delivery string) int {
	deliverySplit := strings.Split(delivery, " ")
	endTime := deliverySplit[3]
	return convertTo24h(endTime)
}

func getStartTime(delivery string) int {
	deliverySplit := strings.Split(delivery, " ")
	endTime := deliverySplit[1]
	return convertTo24h(endTime)
}
