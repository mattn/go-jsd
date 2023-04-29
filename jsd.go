/*
Package jsd provides string similarity metrics/string distance functions based on the Jaro-Winkler similarity.

This package supports unicode/multibyte strings.
*/
package jsd

// StringDistance returns the edit distance between two strings in the range 0 to 1.
// 0 means an exact match and 1 means there is no similarity.
func StringDistance(lhs, rhs string) float64 {
	return Distance([]rune(lhs), []rune(rhs))
}

// StringSimilarity returns the similarity score between two strings in the range 1 to 0.
// 1 means an exact match and 0 means there is no similarity.
func StringSimilarity(lhs, rhs string) float64 {
	return Similarity([]rune(lhs), []rune(rhs))
}

// Distance is similar to StringDistance but takes rune slices instead of strings.
func Distance(lhs, rhs []rune) float64 {
	return 1 - Similarity(lhs, rhs)
}

// Similarity is similar to StringSimilarity but takes rune slices instead of strings.
func Similarity(lhs, rhs []rune) float64 {
	return winklerSim(lhs, rhs)
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

// Jaro similarity
func jaroSim(lhs, rhs []rune) float64 {
	rl1, rl2 := len(lhs), len(rhs)

	if rl1 == 0 || rl2 == 0 {
		return 0
	}
	if rl1 == rl2 && isSame(lhs, rhs) {
		return 1
	}

	dist := max(rl1, rl2)
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
		if !score1[i] {
			continue
		}
		for !score2[j] {
			j++
		}
		if lhs[i] != rhs[j] {
			n++
		}
		j++
	}
	n /= 2

	return (float64(m)/float64(rl1) + float64(m)/float64(rl2) + (float64(m)-n)/float64(m)) / 3.0
}

// Jaro–Winkler similarity
func winklerSim(lhs, rhs []rune) float64 {
	rl1, rl2 := len(lhs), len(rhs)

	lmax := min(rl1, rl2)
	if lmax > 4 {
		lmax = 4 // max length of common prefix
	}

	l := float64(0) // length of common prefix
	for i := 0; i < lmax; i++ {
		if lhs[i] == rhs[i] {
			l++
		}
	}

	simJ := jaroSim(lhs, rhs)
	p := 0.1 // Winkler's prefix scaling factor
	simW := simJ + l*p*(1-simJ)

	return simW
}
