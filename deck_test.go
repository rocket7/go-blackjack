package main

import (
	"os"
	"testing"
)

// t is test handler
//
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 56 {
		t.Errorf("Expected deck length of 56, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades bu got %v ", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs bu got %v ", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	//remove any previous files
	os.Remove("_decktesting")
	//create a new deck
	deck := newDeck()
	//save deck to file
	deck.saveToFile("_decktesting")
	//load deck from file
	loadedDeck := newDeckFromFile("_decktesting")
	//test length of loaded deck
	if len(loadedDeck) != 56 {
		t.Errorf("Expected 56 cards in deck, got %v ", len(deck))
	}
	//remove deck file
	os.Remove("_decktesting")

}
