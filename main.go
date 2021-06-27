package main

import (
	"flag"
	"fmt"
)

func main() {
	tspPointer := flag.String("tsp", "ch130.tsp", "TSP")
	seedPointer := flag.Int64("seed", 0, "Seed")
	flag.Parse()

	problem := readTspFile(*tspPointer, *seedPointer)
	dm := problem.getDistanceMatrix()
	solution := solution{seed: *seedPointer}
	solution.generateRandomSolution(*problem)
	solution.computeCost(dm)
	fmt.Println(solution.cost)
	solution.twoOpt(dm)
	solution.computeCost(dm)
	fmt.Println(solution.cost)
	final := solution.simulatedAnnealing(dm)
	fmt.Println(final.cost)
}
