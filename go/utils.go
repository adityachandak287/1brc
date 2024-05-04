package main

import "math"

// https://realpython.com/python-rounding/#rounding-half-up
func roundHalfUp(n float64, decimals int) float64 {
	mult := math.Pow(10, float64(decimals))
	return math.Floor(n*mult+0.5) / mult
}
