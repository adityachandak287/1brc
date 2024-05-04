package main

import "math"

// https://realpython.com/python-rounding/#rounding-half-up
func roundHalfUp(n float32, decimals int) float32 {
	mult := math.Pow(10, float64(decimals))
	return float32(math.Floor(float64(n)*mult+0.5) / mult)
}
