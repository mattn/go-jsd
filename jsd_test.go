package jsd

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	var tests = []struct {
		name string
		lhs  string
		rhs  string
		want float64
	}{
		{"test001", "こんにちわ世界", "こんにちわ世界", 1},
		{"test002", "こんにちわ世界", "こにゃちわ世界", 0.876},
		{"test003", "こんにちわ世界", "こにゃにゃちわ世界", 0.813},
		{"test004", "こんにちわ世界", "こんばんわ世界", 0.742},
		{"test005", "こんにちわ世界", "こんにちわ", 0.838},
		{"test006", "こんにちわ世界", "こんばんわ", 0.453},
		{"test007", "世界", "こんにちわ", 0},
		{"test008", "こんにちわ世界", "", 0},
	}

	for _, test := range tests {
		got := StringDistance(test.lhs, test.rhs)
		if math.Abs(got-test.want) > 0.001 {
			t.Errorf("%v: want %v but %v: %v vs %v", test.name, test.want, got, test.lhs, test.rhs)
		}
	}
}
