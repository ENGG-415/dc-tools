// Package playingcards implements simple playing card functions including generating
// a deck of cards, shuffling, adding and drawing cards, etc.
//
// Types:
//   - Card: a single playing card with a suit and value
//   - Suit: an enumeration of the four card suits (Clubs, Diamonds, Hearts, Spades)
//   - Deck: a collection of cards with methods for manipulation
//
// Key functions and methods:
//   - Deck.Create: Initialize a full standard deck of 52 cards
//   - Deck.Shuffle: Randomize the order of cards in the deck
//   - Deck.AddCard: Add a single card to the top of the deck
//   - Deck.TakeTopCard: Remove and return the top card from the deck
//   - NumToCardChar: Convert a numeric card value to its character representation
package playingcards
