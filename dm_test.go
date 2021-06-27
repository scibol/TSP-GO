package main

import (
	"math"
	"testing"
)

func TestDistanceMatrix(t *testing.T) {
	p1 := point{1, 2, 2}
	p2 := point{2, 4, 6}
	p3 := point{3, 10, 20}
	p4 := point{4, 20, 20}
	p5 := point{5, 4, 4}
	ps := []point{p1, p2, p3, p4, p5}

	pr := problem{name: "Test", comment: "Test Comment", dimension: 5, points: ps}
	dm := pr.getDistanceMatrix()

	if dm[0][4] != math.Sqrt(8) {
		t.Errorf("Not good! Expected %f, got %f", math.Sqrt(8), dm[0][4])
	}
}

func TestPathCost(t *testing.T) {
	p1 := point{1, 1, 1}
	p2 := point{2, 5, 5}
	p3 := point{3, 2, 7}
	ps := []point{p1, p2, p3}

	pr := problem{name: "Test", comment: "Test Comment", dimension: 3, bestKnown: 0, points: ps}
	s := solution{path: ps}

	dm := pr.getDistanceMatrix()
	s.computeCost(dm)
	expected := math.Sqrt(32) + math.Sqrt(13) + math.Sqrt(37) // distance p1, p2 = sqrt(32), distance p2, p3 = sqrt(13), distance p3, p1 = sqrt(37)
	if s.cost != expected {
		t.Errorf("Not good! Expected %f, got %f", (math.Sqrt(32) + math.Sqrt(13) + math.Sqrt(37)), s.cost)
	}
}
