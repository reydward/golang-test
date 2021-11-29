package main

import (
	"fmt"
	"sync"
)
var retries []retry

func main()  {
	ids := []int {1,20,30,400,50,60}

	retries = RetryOrdersIntegration(ids)

	for _, retry := range retries {
		fmt.Println(retry)
	}
}

type retry struct {
	OrderID      int
	Success		 bool
	ErrorMessage error
}

func RetryOrdersIntegration(ids []int) []retry {
	wg := new(sync.WaitGroup)

	//call RetryOrderIntegration per each id
	for _, id := range ids {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			isSuccessful, err := RetryOrderIntegration(i)

			retries = append(retries, retry{
				OrderID:      i,
				Success:      isSuccessful,
				ErrorMessage: err,
			})

		}(id)
	}
	wg.Wait()
	//return array idorder, boolean, error
	return retries
}

func RetryOrderIntegration(id int) (bool, error) {
	if id == 1 {
		return false, fmt.Errorf("failed to retry integration: %s", "El error")
	}
	return true, nil
}
