package main

import (
	"fmt"
	"math"
)

const C = 299792458
const AUTHOR = "Luis Masuelli"

const (
	G float64 = 6.674 * 0.000000001
	K = 9000000000.0
)

func main() {
	radius := 3.0
	fmt.Printf("Area of a circle of radius %v: %v", radius, math.Pi * radius * radius)

	q1 := 2.0
	var q2 = 3.0
	var distance float64 = 4.0
	fmt.Printf("Attraction between charges %v and %v with distance %v: %v", q1, q2, distance, K * q1 * q2 / (distance * distance))

	m1 := 1.5
	var m2 = 3.0
	var distance2 float64 = 5.0
	fmt.Printf("Gravity between masses %v and %v with distance %v: %v", m1, m2, distance2, G * m1 * m2 / (distance2 * distance2))
}
