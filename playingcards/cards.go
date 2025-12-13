package playingcards

import (
	"log"
	"strconv"
)

// Suit represents the suit of a playing card.
type Suit int

const (
	// Clubs suit
	Clubs Suit = iota
	// Diamonds suit
	Diamonds
	// Hearts suit
	Hearts
	// Spades suit
	Spades
)

// [Card] represents a single playing card.
// Values in [Card] need to be exported to use as argument in RPC.
type Card struct {
	// Val is the card value/rank (1=Ace, 2-10=number, 11=Jack, 12=Queen, 13=King)
	Val int
	// CardSuit is the suit of the card
	CardSuit Suit
}

// NumToCardChar converts a numeric card value to its character representation.
// Valid input values are 1-13, where:
//   - 1 returns "A" (Ace)
//   - 2-10 return the numeric string
//   - 11 returns "J" (Jack)
//   - 12 returns "Q" (Queen)
//   - 13 returns "K" (King)
//
// Invalid values will panic.
func NumToCardChar(val int) string {
	if val > 1 && val < 11 {
		return strconv.Itoa(val)
	} else {
		switch val {
		case 1:
			return "A"
		case 11:
			return "J"
		case 12:
			return "Q"
		case 13:
			return "K"
		default:
			log.Panic("Invalid numerical card rank!")
			return ""
		}
	}
}

// String returns a string representation of the card.
func (c Card) String() string {
	var str string
	str += NumToCardChar(c.Val)
	switch c.CardSuit {
	case Clubs:
		//str += "\u2663"
		str += "\xe2\x99\xa3"
	case Diamonds:
		//str += "\u2666"
		str += "\xe2\x99\xa6"
	case Hearts:
		//str += "\u2665"
		str += "\xe2\x99\xa5"
	case Spades:
		//str += "\u2660"
		str += "\xe2\x99\xa0"
	}
	return str
}
