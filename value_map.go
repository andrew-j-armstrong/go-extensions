package extensions

import (
	"math"
	"math/rand"
)

type ValueMap map[interface{}]float64

func (valueMap *ValueMap) GetTotalValue() float64 {
	var totalValue float64
	for _, value := range *valueMap {
		totalValue += value
	}
	return totalValue
}

func (valueMap *ValueMap) SelectFromWheel() interface{} {
	totalWheelValue := valueMap.GetTotalValue()

	wheelChoice := rand.Float64() * totalWheelValue

	for choice, wheelValue := range *valueMap {
		if wheelValue >= wheelChoice {
			return choice
		}

		wheelChoice -= wheelValue
	}

	return nil
}

func (valueMap *ValueMap) GetKeys() *InterfaceSlice {
	keys := make(InterfaceSlice, len(*valueMap))
	i := 0
	for key := range *valueMap {
		keys[i] = key
		i++
	}
	return &keys
}

func (valueMap *ValueMap) GetBestKey() interface{} {
	bestValue := -math.MaxFloat64
	var bestKey interface{}
	for key, value := range *valueMap {
		if value >= bestValue {
			bestKey = key
			bestValue = value
		}
	}

	return bestKey
}

func (valueMap *ValueMap) Clear() {
	for key := range *valueMap {
		delete(*valueMap, key)
	}
}
