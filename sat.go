package main

import (
	"fmt"
	"time"
)

func main() {
	//make empty arrays to store SAT launching from their respective sites
	SAT := make([]int, 0)
	LS1 := make([]int, 0)
	LS2 := make([]int, 0)
	n := 500 //total number of satellites

	//count is a varible used as a counter for iterating through the 500 sattelites that are stored in SAT
	count := 1

	//Store 500 sattelites in SAT array
	for i := 0; i <= n; i++ {
		SAT = append(SAT, i)
	}

	//Iterate through SAT
	for i := 1; i <= len(SAT)-1; i++ {
		//check the condition while iterating through SAT for the first launch site
		if count <= 4 {
			LS1 = append(LS1, i)
			time.Sleep(2400 * time.Millisecond)
			count++ //increment count until the condition is invalid
			//when count reaches 5th satellite we Launch from the first launch site and make a new array LS1
			if count == 5 {
				fmt.Println("LS1: ", LS1)
				LS1 = make([]int, 0)
			}
			//when count reaches 5th satellite we go into this statement until we reach count <= 8
		} else if count > 4 && count <= 8 {
			LS2 = append(LS2, i)
			time.Sleep(2400 * time.Millisecond)
			count++ //increment count until the condition is invalid
			//when count reaches 9th satellite we Launch from the second launch site and make a new array LS2
			if count == 9 {
				fmt.Println("LS2: ", LS2)
				LS2 = make([]int, 0)
				count = 1 //here we reset the value of count for further sattelites
			}

		}
	}
}
