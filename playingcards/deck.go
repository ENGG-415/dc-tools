package playingcards

import (
	"errors"
	"math/rand"
)

type Deck struct {
	Cards []Card
}

func (deck *Deck) AddCard(c Card) error {
	if c.Val < 1 || c.Val > 13 {
		return errors.New("invalid card")
	}
	deck.Cards = append([]Card{c}, deck.Cards...)
	return nil
}

func (deck *Deck) Reset() error {
	deck.Cards = nil
	return nil
}

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

// method to shuffle the deck
// shuffle (see: https://stackoverflow.com/questions/12264789/shuffle-array-in-go)
// https://stackoverflow.com/questions/12321133/how-to-properly-seed-random-number-generator
func (deck *Deck) Shuffle() error {

	// shuffle cards
	// note: we don't need to seed the random number generator as of Go 1.20
	for i := range deck.Cards {
		j := rand.Intn(i + 1)
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	}
	return nil
}

func (deck Deck) NumCards() int {
	return len(deck.Cards)
}

func (deck Deck) String() (s string) {
	space := ""
	for i := range deck.Cards {
		s = s + space + deck.Cards[i].String()
		space = " "
	}
	return
}

func (deck *Deck) TakeTopCard() Card {
	c := deck.Cards[0]
	deck.Cards = deck.Cards[1:]
	return c
}
