package main

import (
	"fmt"
	"math"
	"math/rand"
)

func annealing(energy float64, newEnergy float64, temperature float64) float64 {
	if newEnergy < energy {
		return 1.0
	}
	return math.Exp((energy - newEnergy) / temperature)
}

func (s *solution) simulatedAnnealing(dm [][]float64) *solution {
	temperature := 100000000000.0
	coolingRate := 0.999
	bestSolution := s
	currentSolution := s
	n := len(s.path)
	for temperature > 1 {
		nextSolution := currentSolution
		a := rand.Intn(n)
		b := rand.Intn(n)
		if a != b {
			nextSolution.path[a], nextSolution.path[b] = nextSolution.path[b], nextSolution.path[a]
		}
		nextSolution.computeCost(dm)
		nextSolution.twoOpt(dm)
		bestSolution.computeCost(dm)
		nextSolution.computeCost(dm)
		if annealing(bestSolution.cost, nextSolution.cost, temperature) > rand.Float64() {
			currentSolution = nextSolution
		}
		if currentSolution.cost < bestSolution.cost {
			fmt.Println("HERE!")
			bestSolution = currentSolution
		}
		temperature *= coolingRate
	}
	return bestSolution
}
