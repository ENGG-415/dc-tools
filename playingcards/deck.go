package playingcards

import (
	"errors"
	"math/rand"
)

// [Deck] represents a collection of playing cards.
type Deck struct {
	// Cards is the slice of cards in the deck
	Cards []Card
}

// AddCard adds a card to the top of the deck.
// Returns an error if the card has an invalid rank value.
func (deck *Deck) AddCard(c Card) error {
	if c.Val < 1 || c.Val > 13 {
		return errors.New("invalid card")
	}
	deck.Cards = append([]Card{c}, deck.Cards...)
	return nil
}

// Reset removes all cards from the deck.
func (deck *Deck) Reset() error {
	deck.Cards = nil
	return nil
}

// Create initializes the deck with a complete standard deck of 52 cards.
// Any existing cards are discarded.
func (deck *Deck) Create() error {
	deck.Cards = nil
	for i := 1; i <= 13; i++ {
		deck.AddCard(Card{i, Clubs})
		deck.AddCard(Card{i, Diamonds})
		deck.AddCard(Card{i, Hearts})
		deck.AddCard(Card{i, Spades})
	}
	return nil
}

// Shuffle randomizes the order of cards in the deck
// See: https://stackoverflow.com/questions/12264789/shuffle-array-in-go
// Note: As of Go 1.20, the random number generator does not need explicit seeding.
func (deck *Deck) Shuffle() error {

	// shuffle cards
	for i := range deck.Cards {
		j := rand.Intn(i + 1)
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	}
	return nil
}

// NumCards returns the number of cards currently in the deck.
func (deck Deck) NumCards() int {
	return len(deck.Cards)
}

// String returns a string representation of the deck, with cards separated by spaces.
func (deck Deck) String() (s string) {
	space := ""
	for i := range deck.Cards {
		s = s + space + deck.Cards[i].String()
		space = " "
	}
	return
}

// TakeTopCard removes and returns the top card from the deck.
func (deck *Deck) TakeTopCard() Card {
	c := deck.Cards[0]
	deck.Cards = deck.Cards[1:]
	return c
}
