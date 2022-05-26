package main

import (
	"encoding/json"
	"fmt"
)

type MultiplyRequest struct {
	Factor1  int32  `json:"factor1, omitempty"`
	Factor2  int32  `json:"factor2, omitempty"`
}
type OdometerInformation struct {
	LastOdometerInformation LastOdometerInformation `json:"lastOdometerInformation,omitempty"`
	ProblemDate             string                  `json:"odometerProblemDate,omitempty"`
	EstimatedCurrentMileage int                     `json:"estimatedCurrentOdometer,omitempty"`
	RecentAnnualMileage     int                     `json:"recentAnnualMileage,omitempty"`
}

type LastOdometerInformation struct {
	Mileage int    `json:"mileage,omitempty"`
	Date    string `json:"date,omitempty"`
}

/*type OdometerInformation struct {
	LastOdometerInformation LastOdometerInformation `json:"lastOdometerInformation,omitempty"`
	ProblemDate        *string `json:"odometerProblemDate"`
	EstimatedCurrentMileage        int `json:"estimatedCurrentOdometer,omitempty"`
	RecentAnnualMileage        int `json:"recentAnnualMileage,omitempty"`
}

type LastOdometerInformation struct {
	Mileage *int `json:"mileage,omitempty"`
	Date *string `json:"date,omitempty"`
}
*/
type ApiPlateResult struct {
	Plate     string   `json:"plate,omitempty"`
	State     string   `json:"state,omitempty"`
	Timestamp string   `json:"timestamp,omitempty"`
	Results   []ApiVin `json:"results"`
}
type ApiVin struct {
	Vin   string `json:"vin,omitempty"`
	Year  int    `json:"year,omitempty"`
	Make  string `json:"make,omitempty"`
	Model string `json:"model,omitempty"`
}

type ApiTLRVTRResult struct {
	Vin        string     `json:"vin"`
	Timestamp  string     `json:"timestamp"`
	Components Components `json:"components"`
}

type Components struct {
	VehicleTitleReport VehicleTitleReport `json:"vehicleTitleReport"`
	TotalLossReport    TotalLossReport    `json:"totalLossReport"`
}

type TotalLossReport struct {
	BodyStyle           string               `json:"bodyStyle"`
	Engine              string               `json:"engine"`
	RecordCount         int                  `json:"recordCount"`
	TitleBrandRecords   []TitleBrandRecords  `json:"titleBrandRecords"`
	TotalLossRecords    []TotalLossRecords   `json:"totalLossRecords"`
	FrameDamageRecords  []FrameDamageRecords `json:"frameDamageRecords"`
	OdometerInformation OdometerInformation  `json:"	odometerInformation"`
}

type TitleBrandRecords struct {
	Brand  string `json:"titleBrands,omitempty"`
	Title  string `json:"titleNumber,omitempty"`
	Date   string `json:"date,omitempty"`
	Source string `json:"source,omitempty"`
	City   string `json:"city,omitempty"`
	State  string `json:"state,omitempty"`
}

type TotalLossRecords struct {
	Date     string `json:"date,omitempty"`
	Source   string `json:"source,omitempty"`
	LossType string `json:"totalLossTypes,omitempty"`
}

type FrameDamageRecords struct {
	Date   string `json:"date,omitempty"`
	Source string `json:"source,omitempty"`
	City   string `json:"city,omitempty"`
	State  string `json:"state,omitempty"`
}

type VehicleTitleReport struct {
	BodyStyle           string               `json:"bodyStyle,omitempty"`
	Engine              string               `json:"engine,omitempty"`
	LastTitleDate       string               `json:"lastTitleDate,omitempty"`
	LastTitleCity       string               `json:"lastTitleCity,omitempty"`
	LastTitleState      string               `json:"lastTitleState,omitempty"`
	Lien                bool                 `json:"lien,omitempty"`
	LienHolderName      string               `json:"lienHolderName,omitempty"`
	RecordCount         int                  `json:"recordCount,omitempty"`
	TitleOrLienRecords  []TitleOrLienRecords `json:"titleOrLienRecords"`
	OdometerInformation OdometerInformation  `json:"	odometerInformation"`
}

type TitleOrLienRecords struct {
	Date         string       `json:"date,omitempty"`
	Mileage      int          `json:"mileage,omitempty"`
	Municipality Municipality `json:"municipality,omitempty"`
}

type Municipality struct {
	Source    string   `json:"source,omitempty"`
	SourceUrl string   `json:"sourceUrl,omitempty"`
	Location  Location `json:"location,omitempty"`
}

type Location struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}


func main()  {
//	var multiplyRequest MultiplyRequest
/*	var odometerInformation OdometerInformation
//	var jsonStr = []byte(`{"factor1": 3, "factor2": 43}`)
	var jsonOdometer = []byte(`{"lastOdometerInformation":{"mileage":5,"date":"2019-03-08"},"odometerProblemDate":null,"estimatedCurrentOdometer":24850}`)
	json.Unmarshal(jsonOdometer, &odometerInformation)
	fmt.Println("odometerInformation", odometerInformation)
*/
/*	var apiPlateResult ApiPlateResult
	var body = []byte(`{"plate":"BYW0326","state":"AZ","timestamp":"2021-01-21T11:58:38.000","results":[{"vin":"ZACCJBBT4FPB89171","year":2015,"make":"JEEP","model":"RENEGADE LATITUDE"}]}`)
	err := json.Unmarshal(body, &apiPlateResult)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("apiPlateResult", apiPlateResult)
*/
	var apiTLRVTRResult ApiTLRVTRResult
	var body = []byte(`{"vin":"JB3BA38K0FU600419","timestamp":"2022-04-04T14:54:06.000","components":{"totalLossReport":{"year":1985,"make":"DODGE","model":"COLT DL","bodyStyle":"Hatchback 4 DR","engine":" 1.5L I4  2BL SOHC","titleBrandRecords":[{"date":"1986-09-10","titleBrands":["Canadian Total Loss"],"source":"Damage Report","city":"","state":"","titleNumber":""}],"totalLossRecords":[{"date":"1986-09-10","source":"Damage Report","totalLossTypes":["Canadian Total Loss"]}],"frameDamageRecords":[],"odometerInformation":{"lastOdometerInformation":{"mileage":null,"date":null},"odometerProblemDate":null,"estimatedCurrentOdometer":null,"recentAnnualMileage":null},"recordCount":1},"vehicleTitleReport":{"year":1985,"make":"DODGE","model":"COLT DL","bodyStyle":"Hatchback 4 DR","engine":" 1.5L I4  2BL SOHC","lastTitleDate":null,"lastTitleCity":null,"lastTitleState":null,"lien":false,"lienHolderName":"","titleOrLienRecords":[],"odometerInformation":{"lastOdometerInformation":{"mileage":null,"date":null},"odometerProblemDate":null,"estimatedCurrentOdometer":null},"recordCount":1}},"customCalculationOdometer":""}`)
	err := json.Unmarshal(body, &apiTLRVTRResult)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("apiTLRVTRResult", apiTLRVTRResult)
}
