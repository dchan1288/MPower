package main

import (
	"log"
)

///
/// question1
/// p1 and p2 are the input parameters
///
func isSubset(p1, p2 []string) bool {
	// c1 each char in p1
	// c2 each char in p2
	subSet := true
	for _, c2 := range p2 {
		charExists := false
		for _, c1 := range p1 {
			if c2 == c1 {
				charExists = true
				break
			}
		}
		subSet = subSet && charExists
	}

	return subSet
}


func main() {
	// Question 1
	log.Println(isSubset([]string{"A", "B", "C", "D", "E"}, []string{"A", "E", "D"}))
	log.Println(isSubset([]string{"A", "B", "C", "D", "E"}, []string{"A", "D", "Z"}))
	log.Println(isSubset([]string{"A", "D", "E"}, []string{"A", "A", "D", "E"}))
	// O(n²)
	// each inner loop will loop through n-1 times of the outer loop
}

