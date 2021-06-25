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
	temperature := 100000.0
	coolingRate := 0.97
	bestSolution := s
	currentSolution := s
	n := len(s.path)
	for temperature > 1 {
		fmt.Println(1)
		nextSolution := bestSolution
		fmt.Println(2)
		a := rand.Intn(n)
		b := rand.Intn(n)
		if a != b {
			nextSolution.path[a], nextSolution.path[b] = nextSolution.path[b], nextSolution.path[a]
		}
		fmt.Println(3)
		nextSolution.twoOpt(dm)
		fmt.Println(4)
		bestSolution.computeCost(dm)
		fmt.Println(5)
		nextSolution.computeCost(dm)
		fmt.Println(6)
		currentSolution.computeCost(dm)
		fmt.Println(7)
		if annealing(bestSolution.cost, nextSolution.cost, temperature) > rand.Float64() {
			currentSolution = nextSolution
			fmt.Println("HERE2!")
		}
		fmt.Println(8)
		if nextSolution.cost < bestSolution.cost {
			fmt.Println("HERE3!")
			bestSolution = currentSolution
		}
		fmt.Println(9)
		temperature *= coolingRate
	}
	fmt.Println("HERE!")
	return bestSolution
}
