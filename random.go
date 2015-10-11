package main

import (
	"math/rand"
)

// randomString returns a random string of letters of the given length
func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		rint := randomInt(65, 117)
		if rint > 90 {
			rint = rint + 6
		}
		bytes[i] = byte(rint)
	}
	return string(bytes)
}

// randomInt returns a random integer between the two numbers.
// min is inclusive, max is exclusive
func randomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
