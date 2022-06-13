package main

import (
	"math"
	"time"
)

// Question 2
type Cache struct {
	timeAccess int64
	Weight     float64
	Value      string
}

var mCache = make(map[string]Cache)
var cacheCapcity = 10
var weight = -1.0

func getCache(key string) interface{} {
	if val, exists := mCache[key]; exists {
		tNow := time.Now().Unix()
		// set weight
		val.Weight /= math.Log(float64((tNow - val.timeAccess + 1)))
		val.timeAccess = tNow
		mCache[key] = val
		return val
	}
	return -1
}

func putCache(key, value string) {
	// check exists
	szCache := len(mCache)
	tNow := time.Now().Unix()
	if szCache >= cacheCapcity { // reach capacity
		// remove the least score
		var leastWeight = math.MaxFloat64
		var theKey string
		for k, v := range mCache {
			if leastWeight > v.Weight {
				leastWeight = v.Weight
				theKey = k
			}
		}
		delete(mCache, theKey)
	}
	// set weight
	mCache[key] = Cache{timeAccess: tNow, Value: value, Weight: weight / -100}
}

// Question 2

func main() {
	// Question 2
	// by using the map structure we have eliminate the used of looping to look up
	// it is as fast as binary lookup O(1)

}
