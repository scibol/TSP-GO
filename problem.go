package main

import "math"

type problem struct {
	name       string
	comment    string
	dimension  int
	best_known int
	points     []point
}

func (pr problem) getDistanceMatrix() [][]float64 {
	dm := make([][]float64, pr.dimension, pr.dimension)

	return dm
}

func (p1 point) distanceFrom(p2 point) float64 {
	diff_x := p1.x - p2.x
	diff_y := p2.y - p2.y
	ss := math.Pow(diff_x, 2) + math.Pow(diff_y, 2)
	return math.Sqrt(ss)

}
