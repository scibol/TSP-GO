package main

import (
	"math"
)

func (s *solution) twoOpt(dm [][]float64) {
	improve := true
	bestGain := 0.0
	localGain := 0.0
	first := 0
	second := 0
	n := len(s.path)

	for improve {
		improve = false
		for i := 0; i < n; i++ {
			bestGain = 0.0
			for j := i + 2; j < n; j++ {
				a := s.path[i]
				b := s.path[i+1]
				c := s.path[j]
				d := s.path[(j+1)%n]

				localGain = moveCost(dm[a.id-1][b.id-1], dm[a.id-1][c.id-1], dm[c.id-1][d.id-1], dm[b.id-1][d.id-1])
				if localGain < bestGain {
					bestGain = localGain
					first = i
					second = j
				}
			}
			if bestGain < -0.00000001 {
				improve = true
				s.swap(first+1, second)
			}
		}
	}
}

func (s *solution) swap(i int, j int) {
	for i < j {
		s.path[i], s.path[j] = s.path[j], s.path[i]
		i++
		j--
	}
}

func moveCost(ab float64, ac float64, cd float64, bd float64) float64 {
	if (ab < ac) && (cd < bd) {
		return 1
	}
	return math.Sqrt(ac) + math.Sqrt(bd) - math.Sqrt(ab) - math.Sqrt(cd)
}
