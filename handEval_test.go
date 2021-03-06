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
// 10 23  36  49   10 Q
// 11 24  37  50   11 K
// 12 25  38  51   12 A

// tests for straight-flush
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
	{ // 2c Ad  5c 4c 3c 6c Qc
		in:  []int{0, 25, 3, 2, 1, 4, 10},
		out: []int{8, 4, 3, 2, 1, 0}, //6-high striaght flush
	},
	{
		in:  []int{1, 23, 51, 48, 50, 49, 47},
		out: []int{8, 12, 11, 10, 9, 8},
	},
}

// tests for four of a kind
var fourOfaKindTests = []struct {
	in  []int
	out []int
}{
	{ // 2c 2d  6h  2h  Th  Ac  2s
		in:  []int{0, 13, 30, 26, 34, 12, 39},
		out: []int{7, 0, 0, 0, 0, 12},
	},
}

// tests for full house
var fullHouseTests = []struct {
	in  []int
	out []int
}{
	{ // 2c 2d Ad 9s Th Ac 2s
		in:  []int{0, 13, 25, 46, 34, 12, 39},
		out: []int{6, 0, 0, 0, 12, 12},
	},
}

// tests for flush
var flushTests = []struct {
	in  []int
	out []int
}{
	{ // 2c 2d 4d 7d 9d Ad 8s
		in:  []int{0, 13, 15, 18, 20, 25, 45},
		out: []int{5, 12, 7, 5, 2, 0},
	},
}

// tests for straight
var straightTests = []struct {
	in  []int
	out []int
}{
	{ // 2c Ad 5c 4c 3d 6h Qc
		in:  []int{0, 25, 3, 2, 14, 30, 10},
		out: []int{4, 4, 3, 2, 1, 0},
	},
	{ // 2c Ad 5c 4c 3d 9h Qc
		in:  []int{0, 25, 3, 2, 14, 33, 10}, // wheel
		out: []int{4, 4, 3, 2, 1, 0},
	},
	{ // Ac Kd Qh Js Th 2s 3s
		in:  []int{12, 24, 36, 48, 34, 39, 40}, // Ace high straight
		out: []int{4, 12, 11, 10, 9, 8},
	},
}

// tests for three of a kind
var setTests = []struct {
	in  []int
	out []int
}{
	{ // 2c 2d 6h 9s Th Ac 2s
		in:  []int{0, 13, 30, 46, 34, 12, 39},
		out: []int{3, 0, 0, 0, 12, 8},
	},
}

// tests for two-pair
var tpTests = []struct {
	in  []int
	out []int
}{
	{
		in:  []int{45, 32, 8, 0, 29, 42, 50},
		out: []int{2, 6, 6, 3, 3, 11},
	},
}

// tests for 1 pair
var pairTests = []struct {
	in  []int
	out []int
}{
	{ // 2c 2d 6h 9s Th Ac Jc
		in:  []int{0, 13, 30, 46, 34, 12, 9},
		out: []int{1, 0, 0, 12, 9, 8},
	},
}

// tests for high card
var highCardTests = []struct {
	in  []int
	out []int
}{
	{ // 2c 4d 6h 9s Th Ac Jc
		in:  []int{0, 15, 30, 46, 34, 12, 9},
		out: []int{0, 12, 9, 8, 7, 4},
	},
}

func TestStraightFlush(t *testing.T) {
	for _, test := range sfTests {
		card := sliceToCards(test.in)
		sort.Sort(ByRank(card))
		sort.Sort(BySuit(card))
		s := [7]Card{}
		copy(s[:], card)

		//t.Log("card", card, test.in)
		if x := straightFlush(s); x == nil || x[0] != test.out[0] {
			t.Errorf("StraightFlush(%v) = %v, want %v", test.in, x, test.out)
		}
	}
}

func TestFourOfaKind(t *testing.T) {
	for _, test := range fourOfaKindTests {
		card := sliceToCards(test.in)
		sort.Sort(ByRank(card))
		s := [7]Card{}
		copy(s[:], card)

		//t.Log("card", card, test.in)
		if x := fourOfaKind(s); x == nil || x[0] != test.out[0] {
			t.Errorf("quads(%v)? got %v, want %v", test.in, x, test.out)
		}
	}
}

func TestFullHouse(t *testing.T) {
	for _, test := range fullHouseTests {
		card := sliceToCards(test.in)
		sort.Sort(ByRank(card))
		s := [7]Card{}
		copy(s[:], card)

		//t.Log("card", card, test.in)
		if x := fullHouse(s); x == nil || x[0] != test.out[0] {
			t.Errorf("fullHouse(%v)? got %v, want %v", test.in, x, test.out)
		}
	}
}

func TestFlush(t *testing.T) {
	for _, test := range flushTests {
		card := sliceToCards(test.in)
		sort.Sort(ByRank(card))
		s := [7]Card{}
		copy(s[:], card)

		//t.Log("card", card, test.in)
		if x := flush(s); x == nil || x[0] != test.out[0] {
			t.Errorf("flush(%v)? got %v, want %v", test.in, x, test.out)
		}
	}
}

func TestStraight(t *testing.T) {
	for _, test := range straightTests {
		card := sliceToCards(test.in)
		sort.Sort(ByRank(card))
		s := [7]Card{}
		copy(s[:], card)

		//t.Log("card", card, test.in)
		if x := straight(s); x == nil || x[0] != test.out[0] {
			t.Errorf("straight(%v)? got %v, want %v", test.in, x, test.out)
		}
	}
}

func TestSet(t *testing.T) {
	for _, test := range setTests {
		card := sliceToCards(test.in)
		sort.Sort(ByRank(card))
		s := [7]Card{}
		copy(s[:], card)

		//t.Log("card", card, test.in)
		if x := set(s); x == nil || x[0] != test.out[0] {
			t.Errorf("set(%v)? got %v, want %v", test.in, x, test.out)
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

func TestPair(t *testing.T) {
	for _, test := range pairTests {
		card := sliceToCards(test.in)
		sort.Sort(ByRank(card))
		s := [7]Card{}
		copy(s[:], card)

		//t.Log("card", card, test.in)
		if x := pair(s); x == nil || x[0] != test.out[0] {
			t.Errorf("pair(%v)? got %v, want %v", test.in, x, test.out)
		}
	}
}

func TestHighCard(t *testing.T) {
	for _, test := range highCardTests {
		card := sliceToCards(test.in)
		sort.Sort(ByRank(card))
		s := [7]Card{}
		copy(s[:], card)

		//t.Log("card", card, test.in)
		if x := highCard(s); x == nil || x[0] != test.out[0] {
			t.Errorf("highCard(%v)? got %v, want %v", test.in, x, test.out)
		}
	}
}
