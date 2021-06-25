package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	id int
	x  float64
	y  float64
}

func readTspFile(tsp string, seed int64) *problem {
	file, err := os.Open("tsp/" + tsp)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pr := problem{}
	for scanner.Scan() {
		line := (strings.Split(scanner.Text(), " "))
		switch line[0] {
		case "NAME:":
			pr.name = line[1]
		case "COMMENT:":
			pr.comment = strings.Join(line[1:], " ")
		case "DIMENSION:":
			pr.dimension, _ = strconv.Atoi(line[1])
		case "BEST_KNOWN":
			pr.bestKnown, _ = strconv.Atoi(line[2])
		default:
			if id, err := strconv.Atoi(line[0]); err == nil {
				x_coord, _ := strconv.ParseFloat(line[1], 64)
				y_coord, _ := strconv.ParseFloat(line[2], 64)
				p := point{id: id, x: x_coord, y: (y_coord)}
				pr.points = append(pr.points, p)
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &pr
}
