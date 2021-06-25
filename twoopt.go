package main

func (s *solution) twoOpt(dm [][]float64) {
	improve := true
	bestGain := 0.0
	totalGain := 0.0
	first := 0
	second := 0
	n := len(s.path)

	for improve {
		for i := 0; i < n; i++ {
			bestGain = 0.0
			for j := i + 2; j < n; j++ {
				a := s.path[i]
				b := s.path[i+1]
				c := s.path[j]
				d := s.path[(j+1)%n]
				if !(dm[a.id-1][b.id-1] > dm[b.id-1][c.id-1]) && !(dm[c.id-1][d.id-1] > dm[c.id-1][a.id-1]) {
					continue
				}
				if b == c || d == a {
					continue
				}
				localGain := -dm[a.id-1][b.id-1] - dm[c.id-1][d.id-1] + dm[a.id-1][c.id-1] + dm[b.id-1][d.id-1]
				if localGain < bestGain {
					bestGain = localGain
					first = i
					second = j
					improve = true
				}
			}
			if bestGain < 0 {
				totalGain += bestGain
				s.path[first], s.path[second] = s.path[second], s.path[first]
			}
		}
	}
}
