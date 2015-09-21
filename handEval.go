package handEval

import (
	"fmt"

	"sort"
	"strconv"
	s "strings"
	//"unicode/utf8"
)

type Card struct {
	rank int
	suit int
}

var p = fmt.Println

type CardSlice [7]Card

func (c CardSlice) String() string {
	// [{7 0} {6 0} {5 0} {4 0} {3 0} {2 0} {1 0}]

	s := ""

	for i := 0; i < len(c); i++ {
		if c[i].rank == 10 {
			s = s + "A"
		} else if c[i].rank == 11 {
			s = s + "B"
		} else if c[i].rank == 12 {
			s = s + "C"
		} else {
			s = s + strconv.Itoa(c[i].rank)
		}
	}
	s = s + "#"
	for i := 0; i < len(c); i++ { // CDHS
		if c[i].suit == 0 {
			s = s + "♣"
		} else if c[i].suit == 1 {
			s = s + "♦"
		} else if c[i].suit == 2 {
			s = s + "♥"
		} else if c[i].suit == 3 {
			s = s + "♠"
		}

	} // ♠♥♦♣

	//fmt.Println(s)
	return s
}

type ByRank []Card

func (c ByRank) Len() int {
	return len(c)
}

func (c ByRank) Less(i, j int) bool {
	return c[i].rank > c[j].rank
}

func (c ByRank) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type BySuit []Card

func (c BySuit) Len() int {
	return len(c)
}

func (c BySuit) Less(i, j int) bool {
	return c[i].suit < c[j].suit
}

func (c BySuit) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

var unshuffled []Card = newDeck()

func newDeck() []Card {
	c := make([]Card, 52)
	k := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			c[k] = Card{j, i}
			k++
		}
	}
	return c
}

func HandEval(cards []int) []int {

	//var int handValue  // the desired result
	c := sliceToCards(cards)

	// fmt.Println("sliceToCards", c)
	//fmt.Println("cards", cards)
	sort.Sort(ByRank(c))
	// fmt.Println("after sorting", c)

	a := [7]Card{}
	copy(a[:], c)

	//fmt.Println("before sorting", c)
	sort.Sort(BySuit(c))
	sf := [7]Card{}
	copy(sf[:], c)
	//fmt.Println("after sorting", s)

	if result := StraightFlush(sf); result != nil {

		return result

	} else if result := FourOfaKind(a); result != nil {

		return result

	} else if result := fullHouse(a); result != nil {

		return result

	} else if result := flush(a); result != nil {

		return result

	} else if result := straight(a); result != nil {

		return result

	} else if result := set(a); result != nil {

		return result

	} else if result := twoPair(a); result != nil {

		return result

	} else if result := pair(a); result != nil {

		return result

	} else if result := highCard(a); result != nil {

		return result

	}

	// panic
	return []int{}
}

// 8
func StraightFlush(c [7]Card) []int {
	//p("from sf", c)

	searchIn := CardSlice(c).String()
	cards := s.Split(searchIn, "#")
	ranks := cards[0]
	suits := cards[1]

	//fmt.Println("cardSplit sf", cards)

	sRanks := []string{"CBA98", "BA987", "A9876", "98765", "87654", "76543",
		"65432", "54321", "43210"}

	// ♠♥♦♣
	sSuits := []string{"♣♣♣♣♣", "♦♦♦♦♦", "♥♥♥♥♥", "♠♠♠♠♠"}

	for i := 0; i < len(sRanks); i++ {
		if index := s.Index(ranks, sRanks[i]); index != -1 {
			// fmt.Println("there is a straight", index)
			for j := 0; j < len(sSuits); j++ {
				// fmt.Println("straight flush", suits[index:], sSuits[j])
				if s.Contains(suits[index:], sSuits[j]) {
					//fmt.Println("there is a straight flush", index, sSuits[j])
					return []int{8, c[index].rank, c[index+1].rank, c[index+2].rank,
						c[index+3].rank, c[index+4].rank}
				}
			}
		}
	}

	// check for a wheel
	wSuits := []string{"♣♣♣♣", "♦♦♦♦", "♥♥♥♥", "♠♠♠♠"} // wheel suits
	if index := s.Index(ranks, "3210"); index != -1 {
		for j := 0; j < len(wSuits); j++ {
			if wSuits[j] == string([]rune(suits)[index:index+4]) {
				//p("partially -", index, wSuits[j])
				// check for an A that is the same suit as 3210
				for m := 0; m < 7; m++ {
					if c[m].rank == 12 && c[m].suit == j {
						//p("wheel straight flush")
						return []int{8, 3, 2, 1, 0}
						break
					}
				}
				break
			}
		}
	}
	return nil
}

// 7
func FourOfaKind(c [7]Card) []int {
	//p("from quads", c)
	searchIn := CardSlice(c).String()
	cards := s.Split(searchIn, "#")
	ranks := cards[0]

	quad := []int{}
	remaining := []Card{}

	sRanks := []string{"CCCC", "BBBB", "AAAA", "9999", "8888", "7777", "6666",
		"5555", "4444", "3333", "2222", "1111", "0000"}

	for i := 0; i < len(sRanks); i++ {
		if index := s.Index(ranks, sRanks[i]); index != -1 {
			//fmt.Println("there is a quad", index)
			quad = []int{7, c[index].rank, c[index+1].rank, c[index+2].rank, c[index+3].rank}

			remaining = append(c[0:index], c[index+4:7]...)

			//p("quads")
			return append(quad, remaining[0].rank)
		}
	}

	return nil
}

// 6
func fullHouse(c [7]Card) []int {
	//p("from boat", c)
	set := []int{6}
	remaining := []Card{}

	// find the set
	for i := 0; i < 5; i++ {
		if c[i].rank == c[i+1].rank && c[i].rank == c[i+2].rank {
			set = append(set, c[i].rank, c[i+1].rank, c[i+2].rank)
			remaining = append(c[0:i], c[i+3:7]...)
		}
	}

	//p("remaining", remaining, set)
	// find the pair
	if len(set) > 1 {
		for i := 0; i < 3; i++ {
			if remaining[i].rank == remaining[i+1].rank {
				//p("full house yo")
				return append(set, remaining[i].rank, remaining[i+1].rank)
			}
		}
		//p("remaining", remaining)
	}

	//p("from end of boat", c)

	return nil
}

// 5
func flush(c [7]Card) []int {

	// searchIn := CardSlice(c).String()
	// cards := s.Split(searchIn, "#")
	// ranks := cards[0]
	// suits := cards[1]

	//p("from flush", c, ranks, suits)

	clubs := make([]int, 1, 6)
	diamonds := make([]int, 1, 6)
	hearts := make([]int, 1, 6)
	spades := make([]int, 1, 6)

	clubs[0] = 5
	diamonds[0] = 5
	hearts[0] = 5
	spades[0] = 5

	for i := 0; i < len(c); i++ {
		//p(clubs, diamonds, hearts, spades)
		if c[i].suit == 0 {
			clubs = append(clubs, c[i].rank)
			if len(clubs) == 6 {
				return clubs
			}
		} else if c[i].suit == 1 {
			diamonds = append(diamonds, c[i].rank)
			if len(diamonds) == 6 {
				return diamonds
			}
		} else if c[i].suit == 2 {
			hearts = append(hearts, c[i].rank)
			if len(hearts) == 6 {
				return hearts
			}
		} else if c[i].suit == 3 {
			spades = append(spades, c[i].rank)
			if len(spades) == 6 {
				return spades
			}
		}
	}

	return nil
}

// 4
func straight(c [7]Card) []int {
	searchIn := CardSlice(c).String()
	//p("searchIn", searchIn)
	cards := s.Split(searchIn, "#")
	ranks := cards[0]
	// highest card of the straight
	h := 12

	sRanks := []string{"CBA98", "BA987", "A9876", "98765", "87654",
		"76543", "65432", "54321", "43210"}

	for i := 0; i < len(sRanks); i++ {
		if index := s.Index(ranks, sRanks[i]); index != -1 {
			//fmt.Println("there is a straight", h, []int{4, h, h - 1, h - 2, h - 3, h - 4})
			return []int{4, h, h - 1, h - 2, h - 3, h - 4}
		}
		h--
	}

	// check for a wheel
	if s.Index(ranks, "3210") != -1 && s.Index(ranks, "C") != -1 {
		//p("wheel")
		return []int{4, 3, 2, 1, 0}
	}

	return nil
}

// 3
func set(c [7]Card) []int {
	searchIn := CardSlice(c).String()
	cards := s.Split(searchIn, "#")
	r := cards[0]
	kickers := make([]int, 0, 2)

	for i := 0; i < 5; i++ {
		if r[i] == r[i+1] && r[i] == r[i+2] {
			setIndex := map[int]int{i: i, i + 1: i + 1, i + 2: i + 2}
			for j := 0; j < 7; j++ {
				if _, ok := setIndex[j]; !ok {
					kickers = append(kickers, j)
				}
				if len(kickers) == 2 {
					break
				}
			}
			// p("there is a set", append([]int{3}, c[i].rank, c[i+1].rank, c[i+2].rank,
			// 	c[kickers[0]].rank, c[kickers[1]].rank))
			return append([]int{3}, c[i].rank, c[i+1].rank, c[i+2].rank,
				c[kickers[0]].rank, c[kickers[1]].rank)
			break
		}
	}

	return nil
}

// 2
func twoPair(c [7]Card) []int {
	searchIn := CardSlice(c).String()
	cards := s.Split(searchIn, "#")
	r := cards[0]
	sl := make([]int, 1, 5)
	// to indicate the hand is two pair
	sl[0] = 2
	pairindex := make(map[int]int)
	kicker := -1

	// find pairs
	for i := 0; i < 6; i++ {
		if r[i] == r[i+1] {
			sl = append(sl, c[i].rank, c[i+1].rank)
			pairindex[i] = i
			pairindex[i+1] = i + 1
			if len(sl) == 5 {
				break
			}
		}
	}
	// find kicker
	// 5 because it includes the indexes for the two pairs and
	// the hand indicator(2)
	if len(sl) == 5 {
		for i := 0; i < 7; i++ {
			if _, ok := pairindex[i]; !ok {
				kicker = i
				break
			}
		}
		hand := append(sl, c[kicker].rank)
		//fmt.Println("there is a two pair", hand)
		return hand
	}
	return nil
}

// 1
func pair(c [7]Card) []int {
	searchIn := CardSlice(c).String()
	cards := s.Split(searchIn, "#")
	r := cards[0]
	pairIndex := -1

	// find the pair
	for i := 0; i < 6; i++ {
		if r[i] == r[i+1] {
			pairIndex = i
			//p("there is a pair", i)
			break
		}
	}
	// prepare return value
	if pairIndex != -1 {
		sl := []int{1, c[pairIndex].rank, c[pairIndex+1].rank}
		count := 0
		rest := make([]int, 3, 3)
		// find the 3 kickers
		for j := 0; j < 7; j++ {
			// done collecting the 3 kickers
			if count == 3 {
				break
			}
			if j == pairIndex || j == pairIndex+1 {
				continue
			} else {
				rest[count] = c[j].rank
				count++
			}
		}
		//p("pair", append(sl, rest[0], rest[1], rest[2]))
		return append(sl, rest[0], rest[1], rest[2])
	}
	//p("not pair")
	return nil
}

// 0
func highCard(c [7]Card) []int {
	sl := make([]int, 6)

	for i := 0; i < 5; i++ {
		sl[i+1] = c[i].rank
	}
	//p("high card", sl)
	return sl
}

func sliceToCards(cards []int) []Card {
	c := make([]Card, 7)
	for i := 0; i < 7; i++ {
		c[i] = unshuffled[cards[i]]
	}

	return c
}

type WinnerReply struct {
	PIndex int
	Type   string
	Hand   []int
}

func Winner(p1 []int, p2 []int) WinnerReply {
	one := HandEval(p1)
	two := HandEval(p2)

	winner := -1

	for index, item := range one {
		if item > two[index] {
			winner = 0
			break
		} else if item < two[index] {
			winner = 1
			break
		}
	}

	if winner == 0 {

		return WinnerReply{0, "", one}

	} else if winner == 1 {

		return WinnerReply{1, "", two}

	} else {

		return WinnerReply{-1, "", []int{}}

	}
}

const (
	CLUBS = iota
	DIAMONDS
	HEARTS
	SPADES
)

func cardValToStr(r int) string {
	switch {
	case r == 0:
		return "2"
	case r == 1:
		return "3"
	case r == 2:
		return "4"
	case r == 3:
		return "5"
	case r == 4:
		return "6"
	case r == 5:
		return "7"
	case r == 6:
		return "8"
	case r == 7:
		return "9"
	case r == 8:
		return "10"
	case r == 9:
		return "J"
	case r == 10:
		return "Q"
	case r == 11:
		return "K"
	case r == 12:
		return "A"
	}
	return fmt.Sprintf("%d", r)
}

func suitToStr(s int) string {
	switch {
	case s == CLUBS:
		return "♣"
	case s == DIAMONDS:
		return "♦"
	case s == HEARTS:
		return "♥"
	case s == SPADES:
		return "♠"
	}
	panic(fmt.Sprintf("invalid suit %d", s))
}

func CardToStr(c []int) string {
	str := ""
	suit := -1

	for _, item := range c {
		if item <= 12 {
			suit = 0
		} else if item <= 25 {
			suit = 1
		} else if item <= 38 {
			suit = 2
		} else if item <= 51 {
			suit = 3
		}
		str += cardValToStr(item%13) + suitToStr(suit) + " "
	}

	return str
}
