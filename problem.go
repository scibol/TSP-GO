package main

import (
	"math"
)

type problem struct {
	name      string
	comment   string
	dimension int
	bestKnown int
	points    []point
}

func (pr problem) getDistanceMatrix() [][]float64 {
	dm := make([][]float64, pr.dimension)
	for i, p1 := range pr.points {
		dm[i] = make([]float64, pr.dimension)
		for j, p2 := range pr.points {
			dm[i][j] = p1.distanceFrom(p2)
		}
	}
	return dm
}

func (p1 point) distanceFrom(p2 point) float64 {
	diff_x := p1.x - p2.x
	diff_y := p2.y - p2.y
	ss := math.Pow(diff_x, 2) + math.Pow(diff_y, 2)
	return math.Sqrt(ss)
}
