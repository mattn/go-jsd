package jsd_test

import (
	"fmt"

	"github.com/mattn/go-jsd"
)

func Example_japanese() {
	for _, test := range []string{
		"こんにちわ世界", "こんちわ世界", "こにゃちわ世界", "こにゃにゃちわ世界",
		"こんばんわ世界", "ハロー・ワールド",
	} {
		lhs := "こんにちわ世界"
		rhs := test
		dist := jsd.StringDistance(lhs, rhs)
		sim := jsd.StringSimilarity(lhs, rhs)

		fmt.Printf("%s <--> %s\n  Dist: %.04f, Sim: %.04f\n", lhs, rhs, dist, sim)
	}
	// Output:
	// こんにちわ世界 <--> こんにちわ世界
	//   Dist: 0.0000, Sim: 1.0000
	// こんにちわ世界 <--> こんちわ世界
	//   Dist: 0.0381, Sim: 0.9619
	// こんにちわ世界 <--> こにゃちわ世界
	//   Dist: 0.0762, Sim: 0.9238
	// こんにちわ世界 <--> こにゃにゃちわ世界
	//   Dist: 0.1429, Sim: 0.8571
	// こんにちわ世界 <--> こんばんわ世界
	//   Dist: 0.1524, Sim: 0.8476
	// こんにちわ世界 <--> ハロー・ワールド
	//   Dist: 1.0000, Sim: 0.0000
}

func Example_english() {
	for _, test := range []string{
		"accomodate", "accommodate", "accommodated", "accommodates", "accommodating",
		"accommodation", "ping ping",
	} {
		lhs := "accomodate"
		rhs := test
		dist := jsd.StringDistance(lhs, rhs)
		sim := jsd.StringSimilarity(lhs, rhs)

		fmt.Printf("%s <--> %s\n  Dist: %.04f, Sim: %.04f\n", lhs, rhs, dist, sim)
	}
	// Output:
	// accomodate <--> accomodate
	//   Dist: 0.0000, Sim: 1.0000
	// accomodate <--> accommodate
	//   Dist: 0.0182, Sim: 0.9818
	// accomodate <--> accommodated
	//   Dist: 0.0333, Sim: 0.9667
	// accomodate <--> accommodates
	//   Dist: 0.0333, Sim: 0.9667
	// accomodate <--> accommodating
	//   Dist: 0.0815, Sim: 0.9185
	// accomodate <--> accommodation
	//   Dist: 0.0815, Sim: 0.9185
	// accomodate <--> ping ping
	//   Dist: 1.0000, Sim: 0.0000
}
