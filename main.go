package main

import (
	"flag"
	"fmt"
)

func main() {
	tspPointer := flag.String("tsp", "ch130.tsp", "TSP")
	seedPointer := flag.Int64("seed", 0, "Seed")
	flag.Parse()

	// fmt.Println(*tspPointer)
	// fmt.Print(*seedPointer)
	problem := readTspFile(*tspPointer, *seedPointer)
	fmt.Println(problem)
}
