package main

import (
	"errors"
	"log"
)

// Question 3
func linear(nTime int) (float64, error) {
	if nTime < 2 {
		return 0, errors.New("invalid input")
	}
	cur := 0.0
	for n := nTime; n >= 2; n-- {
		if n == 2 {
			cur = 1.0 / (float64(n) + cur)
			break
		}
		cur = (cur + 1.0) / float64((n * (n - 1)))
	}
	return cur, nil
}

func main() {
	val, err := linear(10)
	log.Println(val, err)
}
