package xmanipulations

import (
	"math/rand"
)

// Shuffle shuffles the array using a random source
func Shuffle(array []interface{}, source rand.Source) {
	random := rand.New(source)
	for i := len(array) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}
