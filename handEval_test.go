package handEval

import "testing"
import "sort"

// c  d   h   s
// 0  13  26  39   0  2
// 1  14  27  40   1  3
// 2  15  28  41   2  4
// 3  16  29  42   3  5
// 4  17  30  43   4  6
// 5  18  31  44   5  7
// 6  19  32  45   6  8
// 7  20  33  46   7  9
// 8  21  34  47   8  T
// 9  22  35  48   9  J
// 10 23  36  49   A  Q
// 11 24  37  50   B  K
// 12 25  38  51   C  A

var sfTests = []struct {
	in  []int
	out []int
}{
	{
		in:  []int{10, 12, 11, 7, 9, 8, 28},
		out: []int{8, 12, 11, 10, 9, 8},
	},
	{
		in:  []int{12, 25, 3, 2, 1, 24, 0}, //sf wheel
		out: []int{8, 3, 2, 1, 0},
	},
	{
		in:  []int{38, 49, 22, 29, 27, 26, 28}, //sf wheel
		out: []int{8, 3, 2, 1, 0},
	},
	{
		in:  []int{0, 25, 3, 2, 1, 4, 10},
		out: []int{8, 4, 3, 2, 1, 0},
	},
	{
		in:  []int{1, 23, 51, 48, 50, 49, 47},
		out: []int{8, 12, 11, 10, 9, 8},
	},
}

var tpTests = []struct {
	in  []int
	out []int
}{
	{
		in:  []int{45, 32, 8, 0, 29, 42, 50},
		out: []int{2, 6, 6, 3, 3, 11},
	},
}

// wheel
// p(handEval.HandEval([]int{0, 25, 3, 2, 14, 33, 10}))

// straight
// p(handEval.HandEval([]int{0, 25, 3, 2, 14, 30, 10}))

// // Ace high straight
// p(handEval.HandEval([]int{12, 24, 36, 48, 34, 39, 40}))

// // high card
// p(handEval.HandEval([]int{0, 15, 30, 46, 34, 12, 9}))

// // pair
// p(handEval.HandEval([]int{0, 13, 30, 46, 34, 12, 9}))

// // flush
// p(handEval.HandEval([]int{0, 13, 15, 18, 20, 25, 45}))

// // set
// p(handEval.HandEval([]int{0, 13, 30, 46, 34, 12, 39}))

// // quads
// p(handEval.HandEval([]int{0, 13, 30, 26, 34, 12, 39}))

// // fullhouse
// p(handEval.HandEval([]int{0, 13, 25, 46, 34, 12, 39}))

func TestStraightFlush(t *testing.T) {
	for _, test := range sfTests {
		card := sliceToCards(test.in)
		sort.Sort(ByRank(card))
		sort.Sort(BySuit(card))
		s := [7]Card{}
		copy(s[:], card)

		//t.Log("card", card, test.in)
		if x := StraightFlush(s); x == nil || x[0] != test.out[0] {
			t.Errorf("StraightFlush(%v) = %v, want %v", test.in, x, test.out)
		}
	}
}

func TestTwoPair(t *testing.T) {
	for _, test := range tpTests {
		card := sliceToCards(test.in)
		sort.Sort(ByRank(card))
		s := [7]Card{}
		copy(s[:], card)

		//t.Log("card", card, test.in)
		if x := twoPair(s); x == nil || x[0] != test.out[0] {
			t.Errorf("twoPair(%v) = %v, want %v", test.in, x, test.out)
		}
	}
}
