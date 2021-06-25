package main

import (
	"math/rand"
)

type solution struct {
	path []point
	seed int64
	cost float64
}

func (s *solution) generateRandomSolution(pr problem) {
	path := make([]point, pr.dimension)
	for i, value := range rand.Perm(pr.dimension) {
		path[i] = pr.points[value]
	}
	s.path = path
}

func (s *solution) computeCost(dm [][]float64) {
	sum := 0.0
	for i := range s.path {
		p1 := s.path[i]
		p2 := s.path[(i+1)%len(s.path)]
		sum += dm[p1.id-1][p2.id-1]
	}
	s.cost = sum
}
