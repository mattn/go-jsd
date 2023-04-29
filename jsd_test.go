package jsd

import (
	"math"
	"testing"

	"golang.org/x/text/width"
)

var tests = []struct {
	name string
	lhs  string
	rhs  string
	want float64
}{
	{"test001", "こんにちわ世界", "こんにちわ世界", 0},
	{"test002", "こんにちわ世界", "こにゃちわ世界", 0.0761},
	{"test003", "こんにちわ世界", "こにゃにゃちわ世界", 0.1428},
	{"test004", "こんにちわ世界", "こんばんわ世界", 0.1523},
	{"test005", "こんにちわ世界", "こんにちわ", 0.0571},
	{"test006", "こんにちわ世界", "こんばんわ", 0.2590},
	{"test007", "世界", "こんにちわ", 1},
	{"test008", "こんにちわ世界", "", 1},
	{"test009", "", "こんにちわ世界", 1},
	// Test values taken from Rosetta Code:
	// https://rosettacode.org/wiki/Jaro-Winkler_distance
	{"test010", "accomodate", "accommodate", 0.0182},
	{"test011", "accomodate", "accommodated", 0.0333},
	{"test012", "accomodate", "accommodates", 0.0333},
	{"test013", "accomodate", "accommodating", 0.0815},
	{"test014", "accomodate", "accommodation", 0.0815},
	{"test015", "definately", "definitely", 0.0400},
	{"test016", "definately", "defiantly", 0.0422},
	{"test017", "definately", "define", 0.0800},
	{"test018", "definately", "definite", 0.0850},
	{"test019", "definately", "definable", 0.0872},
	{"test020", "goverment", "government", 0.0533},
	{"test021", "goverment", "govern", 0.0667},
	{"test022", "goverment", "governments", 0.0697},
	{"test023", "goverment", "movement", 0.0810},
	{"test024", "goverment", "governmental", 0.0833},
}

func TestStringDistance(t *testing.T) {
	for _, test := range tests {
		want := test.want

		t.Run(test.name, func(t *testing.T) {
			got := StringDistance(test.lhs, test.rhs)
			if math.Abs(got-want) > 0.0001 {
				t.Errorf("%v: want %v but got %v: %v vs %v",
					test.name, test.want, got, test.lhs, test.rhs)
			}
		})

		// Make sure that the same value is returned when converting to full-width
		// characters (eg. "abc" --> "ａｂｃ").
		t.Run(test.name+":narrow", func(t *testing.T) {
			got := StringDistance(
				width.Narrow.String(test.lhs),
				width.Narrow.String(test.rhs),
			)
			if math.Abs(got-want) > 0.0001 {
				t.Errorf("%v: want %v but got %v: %v vs %v",
					test.name, test.want, got, test.lhs, test.rhs)
			}
		})
	}
}

func Test_jaroSim(t *testing.T) {
	lhs := []rune("FAREMVIEL")
	rhs := []rune("FARMVILLE")

	// (1. / 3.) * ((8. / 9.) + (8. / 9.) + ((8. - 1.) / 8.)) = 0.88425925925
	want := 0.884
	got := jaroSim(lhs, rhs)

	if math.Abs(got-want) > 0.001 {
		t.Errorf("want %v but got %v: %v vs %v", want, got, lhs, rhs)
	}
}

func BenchmarkDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			_ = StringDistance(test.lhs, test.rhs)
		}
	}
}
