package misc

import (
	"fmt"
	"math/rand"
	"sort"
)

/*
This problem was originally a python problem as follows:

Class DeckOfCards(object):

	def __init__(self):
		ranks = ['Ace', 'One', 'Two', 'Three', 'Four', 'Five', 'Six', 'Seven', 'Eight', 'Nine', 'Jack', 'Queen', 'King']
		suits = ['Clubs', 'Hearts', 'Diamond', 'Spades']
		# Assemble all the cards here, like 'Ace of Spades'. Both 'One of Diamond' and '1 of Diamond' are valid
		self.deck = []

	def get_cards(self, num_of_cards=5):
		'''Remove and return Top N cards from the deck'''

	def put_back_card(self, card):
		'''Put the provide card at the back of the deck'''

	def shuffle(self):
	  '''Shuffle the deck'''

	def sort(self):
		'''Sort the deck'''
*/

var RANK = []string{
	"Ace",
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Jack",
	"Queen",
	"King",
}

var SUITE = []string{
	"Club",
	"Heart",
	"Diamond",
	"Spade",
}

type Card struct {
	RankOrder  int
	SuiteOrder int
	Card       string
}

type Deck struct {
	Cards []Card
}

func NewCard(rank, suit int, card string) Card {
	return Card{rank, suit, card}
}

func NewDeck() *Deck {
	deck := &Deck{}
	for i, suite := range SUITE {
		for j, rank := range RANK {
			card := NewCard(j, i, rank+" of "+suite)
			deck.Cards = append(deck.Cards, card)
		}
	}
	return deck
}

func (d *Deck) Print() {
	for _, card := range d.Cards {
		fmt.Println(card.Card)
	}
}

func (d *Deck) GetCards(n int) (topCards []Card) {
	// Overall both these will look at different views of the same underlying array
	topCards = d.Cards[:n] // topCards has a view of the first 5 cards of the underlying array of slice Cards
	d.Cards = d.Cards[n:]  // Cards slice now has a shrunken view of its self
	return
}

func (d *Deck) PutBackCard(c Card) {
	d.Cards = append(d.Cards, c)
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) Sort() {
	// Magic sauce: https://stackoverflow.com/a/46200067/768020
	// The idea is same as the sorting in bankingsystem
	sort.Slice(d.Cards, func(i, j int) bool {
		if d.Cards[i].SuiteOrder == d.Cards[j].SuiteOrder {
			return d.Cards[i].RankOrder < d.Cards[j].RankOrder
		} else {
			return d.Cards[i].SuiteOrder < d.Cards[j].SuiteOrder
		}
	})
}
