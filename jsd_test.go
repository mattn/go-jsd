package jsd

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	var tests = []struct {
		lhs  string
		rhs  string
		want float64
	}{
		{"こんにちわ世界", "こんにちわ世界", 1},
		{"こんにちわ世界", "こにゃちわ世界", 0.904},
		{"こんにちわ世界", "こにゃにゃちわ世界", 0.841},
		{"こんにちわ世界", "こんばんわ世界", 0.809},
		{"こんにちわ世界", "こんにちわ", 0.904},
		{"こんにちわ世界", "こんばんわ", 0.676},
		{"こんにちわ世界", "世界", 0},
	}

	for _, test := range tests {
		got := StringDistance(test.lhs, test.rhs)
		if math.Abs(got-test.want) > 0.001 {
			t.Fatalf("want %v but %v: %v vs %v", test.want, got, test.lhs, test.rhs)
		}
	}
}
