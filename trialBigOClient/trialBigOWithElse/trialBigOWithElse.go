package trialBigOWithElse

import (
	"fmt"
	"strconv"
	"time"
)

// This function will return a string array of length 10 with numbers from 1 to 10
func GetListUpto10WithElse() [10]string {
	startTime := time.Now()
	returnVal := [10]string{}
	for i := 0; i < 10; i++ {
		if (i+1)%3 == 0 {
			returnVal[i] = "fizz"

		} else {
			returnVal[i] = strconv.Itoa(i + 1)
		}
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("\n Elapsed time WithElse %d NanoSeconds", elapsedTime.Nanoseconds())
	return returnVal
}
