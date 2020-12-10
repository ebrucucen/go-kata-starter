package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlwaysTrue(t *testing.T) {

	assert.True(t, true)
}

// we want to test when a list of strings up to 10, then it should replace numbers when divisible by 3 with fizz, or 5 buzz , it should ...
// we want to test the function output a list of strings up 10. 
func TestReturnsListOf10Items(t *testing.T) {
	expectedNumberOfItems:=10
	assert.Equal(t, expectedNumberOfItems , len(GetListUpto10()))
}

// This test will verify the content of the array returned has the right values
func TestReturnedItemsAre1to10(t *testing.T) {
	expectedItems:= [10]string{"1", "2", "fizz", "4", "5", "fizz", "7", "8", "fizz", "10"}
	assert.Equal(t, expectedItems, GetListUpto10())
}

// This function will return a string array of length 10 with numbers from 1 to 10
func GetListUpto10() [10]string {
	returnVal:=[10]string{}
		for i:=0; i<10;i++ {
			if (i+1) % 3 ==0 {
				returnVal[i]= "fizz"
				
			}else {
				returnVal[i]= strconv.Itoa(i+1)
			}
	} 
	return returnVal
}


