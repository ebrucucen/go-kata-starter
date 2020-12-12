package trialBigOWithoutElse

import (
	"fmt"
	"strconv"
	"time"
)

// This function will return a string array of length 10 with numbers from 1 to 10
func GetListUpto10WithoutElse() [10]string {
	startTime := time.Now()
	returnVal := [10]string{}
	for i := 0; i < 10; i++ {
		returnVal[i] = strconv.Itoa(i + 1)
		if (i+1)%3 == 0 {
			returnVal[i] = "fizz"
		}
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("\n Elapsed time WithoutElse %d NanoSeconds\n", elapsedTime.Nanoseconds())
	return returnVal
}

func main() {
	GetListUpto10WithoutElse()
}
