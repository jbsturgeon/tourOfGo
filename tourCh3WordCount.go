//Implement WordCount. It should return a map of the counts of each “word” in the string s.
//The wc.Test function runs a test suite against the provided function and prints success
//failure.
//
//You might find strings.Fields helpful.

package main

import (
	_ "fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	sMap := make(map[string]int) //or:  var sMap map[string]int
	var sArray []string = strings.Fields(s)

	for _, iStr := range sArray {
		elem, ok := sMap[iStr]
		if ok {
			sMap[iStr] = elem + 1
		} else {
			sMap[iStr] = 1
		}
	}

	return sMap
}

func main() {
	wc.Test(WordCount)
}
