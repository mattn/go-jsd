package jsd

func StringDistance(lhs, rhs string) float64 {
	return Distance([]rune(lhs), []rune(rhs))
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func isSame(lhs, rhs []rune) bool {
	for i := 0; i < len(lhs); i++ {
		if lhs[i] != rhs[i] {
			return false
		}
	}
	return true
}

func Distance(lhs, rhs []rune) float64 {
	rl1, rl2 := len(lhs), len(rhs)

	if rl1 == 0 || rl2 == 0 {
		return 0
	}
	if rl1 == rl2 && isSame(lhs, rhs) {
		return 1
	}

	dist := rl1
	if rl2 > dist {
		dist = rl2
	}
	dist = dist/2 - 1

	m := 0
	ml := max(rl1, rl2)
	score1 := make([]bool, ml)
	score2 := make([]bool, ml)
	for i := 0; i < rl1; i++ {
		lo := max(0, i-dist)
		hi := min(i+dist+1, rl2)
		for j := lo; j < hi; j++ {
			if score2[j] || lhs[i] != rhs[j] {
				continue
			}
			score1[i] = true
			score2[j] = true
			m++
			break
		}
	}
	if m == 0 {
		return 0
	}

	n := float64(0)
	j := 0
	for i := 0; i < rl1; i++ {
		if score1[i] != score2[j] {
			n++
		}
	}
	n /= 2

	return (float64(m)/float64(rl1) + float64(m)/float64(rl2) + (float64(m)-n)/float64(m)) / 3.0
}
