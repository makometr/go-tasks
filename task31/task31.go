package task31

import "math"

// Point2D has 2D coordinates
type Point2D struct {
	x, y float64
}

// NewPoint creates new point struct
func NewPoint(x, y float64) *Point2D {
	return &Point2D{x, y}
}

// GetDistance calculates distance between 2 2D-Points
func GetDistance(a, b Point2D) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}
