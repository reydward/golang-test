package main

import (
	"encoding/json"
	"fmt"
)

type MultiplyRequest struct {
	Factor1  int32  `json:"factor1, omitempty"`
	Factor2  int32  `json:"factor2, omitempty"`
}

func main()  {
	var multiplyRequest MultiplyRequest
	var jsonStr = []byte(`{"factor1": 3, "factor2": 43}`)
	json.Unmarshal(jsonStr, &multiplyRequest)
	fmt.Println("multiplyRequest", multiplyRequest)
}
