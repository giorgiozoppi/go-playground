package main

import (
	"fmt"
	"math"
)

func computeRoot(x float64, n int) float64 {
	if x == float64(0) {
		return 0
	}
	lowerBound := float64(0)
	upperBound := math.Max(1, x)
	approxRoot := (upperBound - lowerBound) / 2
	for approxRoot-lowerBound >= 0.001 {

		if math.Pow(approxRoot, float64(n)) > x {
			upperBound = approxRoot
		} else if math.Pow(approxRoot, float64(n)) < x {
			lowerBound = approxRoot
		} else {
			break
		}
		approxRoot = (upperBound + lowerBound) / 2
	}
	return approxRoot
}
func main() {
	computeRoot := computeRoot(81, 9)
	fmt.Printf("%f\n", computeRoot)
}
