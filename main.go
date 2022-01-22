package main

import "fmt"

var zzz int // variable declaration is valid
//card := "Ace of Spades" # this is invalid

func main() {
	var card string = "Ace of Spades"
	//card_array := []string{card, "Queen of Spades", newCard()}
	card_array := deck{card, "Queen of Spades", newCard()}
	card_array = append(card_array, "Six of Clubs") //Tis Retuns a NEW Array
	// use := only for new variable declaration

	card = "Seven of Diamonds"
	fmt.Println(card)

	//must use all declared variables
	for i, x := range card_array {
		fmt.Println(i, x)
	}

	// CALLING DECK FUNCTION PRINT
	card_array.print()

	// CALLING NEWDECK FUNCTION PRINT
	cards := newDeck()
	cards.print()
	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	//Return multiple values - 5 cards and remaining cards in deck
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

}

func newCard() string {
	return "Five of Diamonds"

}
