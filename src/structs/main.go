package main

import "fmt"

// Define high cards and suits
const (
	J       = 11
	Q       = 12
	K       = 13
	A       = 14
	CLUB    = 10
	DIAMOND = 20
	HEART   = 30
	SPADE   = 40
)

// HANDSIZE A Poker hand size
const HANDSIZE = 5

// Card is defined by rank and suit
type Card struct {
	rank int
	suit int
}

// PokerHand is defined a set of cards
type PokerHand struct {
	hand [HANDSIZE]Card
}

// NewPokerHand : constructs a Poker Hand object fromn a given set of 5 cards. The poker hand is sorted by rank
func NewPokerHand(cards [HANDSIZE]Card) PokerHand {
	ph := PokerHand{cards}
	ph.sortHandByRank()
	return ph
}

func (ph PokerHand) isHighCard() bool {
	var r = false
	if !ph.isRoyalFlush() &&
		!ph.isFlush() &&
		!ph.isStraight() &&
		!ph.isFullHouse() &&
		!ph.isStraightFlush() &&
		!ph.isThreeOfAKind() &&
		!ph.isFourOfAKind() &&
		!ph.isTwoPairs() &&
		!ph.isOnePair() {
		r = true
	}
	return r
}

// w.w.x.y.z || x.w.w.y.z || x.y.w.w.z || x.y.z.w.w
func (ph PokerHand) isOnePair() bool {
	var r bool

	if ph.isThreeOfAKind() || ph.isFourOfAKind() || ph.isFullHouse() || ph.isTwoPairs() {
		r = false
	} else {
		r = ph.hand[0].rank == ph.hand[1].rank || ph.hand[1].rank == ph.hand[2].rank ||
			ph.hand[2].rank == ph.hand[3].rank || ph.hand[3].rank == ph.hand[4].rank
	}
	return r
}

// x.x.y.y.z || x.y.y.z.z || x.x.y.z.z
func (ph PokerHand) isTwoPairs() bool {
	var r bool

	if ph.isFourOfAKind() || ph.isFullHouse() {
		r = false
	} else {
		r = (ph.hand[0].rank == ph.hand[1].rank && ph.hand[2].rank == ph.hand[3].rank) ||
			(ph.hand[1].rank == ph.hand[2].rank && ph.hand[3].rank == ph.hand[4].rank) ||
			(ph.hand[0].rank == ph.hand[1].rank && ph.hand[3].rank == ph.hand[4].rank)
	}
	return r
}

// x.x.x.y.z || y.x.x.x.z || y.z.x.x.x
func (ph PokerHand) isThreeOfAKind() bool {
	var r bool

	if ph.isFourOfAKind() || ph.isFullHouse() {
		r = false
	} else if (ph.hand[0].rank == ph.hand[1].rank && ph.hand[1].rank == ph.hand[2].rank) ||
		(ph.hand[1].rank == ph.hand[2].rank && ph.hand[2].rank == ph.hand[3].rank) ||
		(ph.hand[3].rank == ph.hand[2].rank && ph.hand[3].rank == ph.hand[4].rank) {
		r = true
	}
	return r
}

//  x.x.x.x.y || y.x.x.x.x
func (ph PokerHand) isFourOfAKind() bool {
	var r1, r2 bool

	r1 = (ph.hand[0].rank == ph.hand[1].rank) && (ph.hand[1].rank == ph.hand[2].rank) && (ph.hand[2].rank == ph.hand[3].rank)
	r2 = (ph.hand[1].rank == ph.hand[2].rank) && (ph.hand[2].rank == ph.hand[3].rank) && (ph.hand[3].rank == ph.hand[4].rank)

	return (r1 || r2)
}

// same as a straight flush but with highest card being an Ace
func (ph PokerHand) isRoyalFlush() bool {
	var r1, r2 bool
	r1 = ph.isStraightFlush()
	r2 = (ph.hand[HANDSIZE-1].rank == A)
	return (r1 && r2)
}

func (ph PokerHand) isStraightFlush() bool {
	return (ph.isFlush() && ph.isStraight())
}

// x.x.y.y.y || x.x.x.y.y
func (ph PokerHand) isFullHouse() bool {
	var r1, r2 bool

	r1 = (ph.hand[0].rank == ph.hand[1].rank) && (ph.hand[2].rank == ph.hand[3].rank) && (ph.hand[3].rank == ph.hand[4].rank)
	r2 = (ph.hand[0].rank == ph.hand[1].rank) && (ph.hand[1].rank == ph.hand[2].rank) && (ph.hand[3].rank == ph.hand[4].rank)

	return (r1 || r2)
}

func (ph PokerHand) isFlush() bool {
	var r bool
	r = (ph.hand[0].suit == ph.hand[1].suit && ph.hand[1].suit == ph.hand[2].suit &&
		ph.hand[2].suit == ph.hand[3].suit && ph.hand[3].suit == ph.hand[4].suit)
	return r
}

func (ph PokerHand) isStraight() bool {
	var r = true
	for i := 0; i < HANDSIZE-1; i++ {
		if ph.hand[i].rank != (ph.hand[i+1].rank - 1) {
			r = false
			break
		}
	}
	return r
}

func (ph *PokerHand) sortHandByRank() {

	var smallest int

	for i := 0; i < HANDSIZE-1; i++ {
		smallest = i
		for j := i + 1; j < HANDSIZE; j++ {
			if ph.hand[j].rank < ph.hand[smallest].rank {
				smallest = j
			}
		}
		c := ph.hand[i]
		ph.hand[i] = ph.hand[smallest]
		ph.hand[smallest] = c
	}
}

func (ph PokerHand) evaluate() string {
	var r string
	if ph.isHighCard() {
		r = "High Card"
	} else if ph.isOnePair() {
		r = "One Pair"
	} else if ph.isTwoPairs() {
		r = "Two Pairs"
	} else if ph.isThreeOfAKind() {
		r = "Three Of A Kind"
	} else if ph.isFourOfAKind() {
		r = "Four Of A Kind"
	} else if ph.isRoyalFlush() {
		r = "Royal Flush"
	} else if ph.isStraightFlush() {
		r = "Straight Flush"
	} else if ph.isFlush() {
		r = "Flush"
	} else if ph.isStraight() {
		r = "Straight"
	} else if ph.isFullHouse() {
		r = "Full House"
	} else {
		r = "Unknown"
	}
	return r
}

func main() {
	cardsHighCard := [5]Card{{2, DIAMOND}, {A, HEART}, {3, CLUB}, {J, HEART}, {10, SPADE}}
	cardsOnePair := [5]Card{{A, DIAMOND}, {2, HEART}, {A, CLUB}, {J, HEART}, {10, SPADE}}
	cardsTwoPairs := [5]Card{{2, DIAMOND}, {2, HEART}, {A, CLUB}, {J, HEART}, {A, SPADE}}
	cardsThreeOfAKind := [5]Card{{2, DIAMOND}, {2, HEART}, {2, CLUB}, {J, HEART}, {10, SPADE}}
	cardsFourOfAKind := [5]Card{{2, DIAMOND}, {2, HEART}, {2, CLUB}, {J, HEART}, {2, SPADE}}
	cardsFullHouse := [5]Card{{J, DIAMOND}, {2, HEART}, {2, CLUB}, {J, HEART}, {2, SPADE}}
	cardsFlush := [5]Card{{3, DIAMOND}, {2, DIAMOND}, {6, DIAMOND}, {5, DIAMOND}, {4, DIAMOND}}
	cardsRoyalFlush := [5]Card{{J, DIAMOND}, {Q, DIAMOND}, {A, DIAMOND}, {10, DIAMOND}, {K, DIAMOND}}
	cardsStraight := [5]Card{{2, DIAMOND}, {3, CLUB}, {4, HEART}, {5, HEART}, {6, SPADE}}

	pokerHand := NewPokerHand(cardsFlush)
	fmt.Println(pokerHand.hand, "=> ", pokerHand.evaluate())

	pokerHand = NewPokerHand(cardsRoyalFlush)
	fmt.Println(pokerHand.hand, "=> ", pokerHand.evaluate())

	pokerHand = NewPokerHand(cardsHighCard)
	fmt.Println(pokerHand.hand, "=> ", pokerHand.evaluate())

	pokerHand = NewPokerHand(cardsOnePair)
	fmt.Println(pokerHand.hand, "=> ", pokerHand.evaluate())

	pokerHand = NewPokerHand(cardsTwoPairs)
	fmt.Println(pokerHand.hand, "=> ", pokerHand.evaluate())

	pokerHand = NewPokerHand(cardsThreeOfAKind)
	fmt.Println(pokerHand.hand, "=> ", pokerHand.evaluate())

	pokerHand = NewPokerHand(cardsFourOfAKind)
	fmt.Println(pokerHand.hand, "=> ", pokerHand.evaluate())

	pokerHand = NewPokerHand(cardsFullHouse)
	fmt.Println(pokerHand.hand, "=> ", pokerHand.evaluate())

	pokerHand = NewPokerHand(cardsStraight)
	fmt.Println(pokerHand.hand, "=> ", pokerHand.evaluate())

}
