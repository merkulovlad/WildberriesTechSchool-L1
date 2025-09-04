package main

import (
	"fmt"
	"math"
)

func groupTemperatures(temps []float64) map[float64][]float64 {
	groups := make(map[float64][]float64)

	for _, t := range temps {
		var key float64
		if t < 0 {
			key = math.Ceil(t/10) * 10
		} else {
			key = math.Floor(t/10) * 10
		}
		groups[key] = append(groups[key], t)
	}

	return groups
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := groupTemperatures(temps)

	for key, values := range groups {
		fmt.Printf("%v: %v\n", key, values)
	}
}
