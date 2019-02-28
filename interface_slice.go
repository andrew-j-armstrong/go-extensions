package extensions

import (
	"math"
	"math/rand"
)

type InterfaceSlice []interface{}

func (slice *InterfaceSlice) GetBestEntry(getValue func(interface{}) float64) interface{} {
	bestValue := -math.MaxFloat64
	var bestEntry interface{}
	for _, entry := range *slice {
		value := getValue(entry)
		if value >= bestValue {
			bestEntry = entry
			bestValue = value
		}
	}

	return bestEntry
}

func (slice *InterfaceSlice) SelectRandom() interface{} {
	return (*slice)[rand.Intn(len(*slice))]
}
