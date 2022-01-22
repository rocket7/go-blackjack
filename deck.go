package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// GO IS NOT OBJECT ORIENTED

//CREATE A NEW "TYPE" CALLED "DECK" WHICH "EXTENDS" SLICE OF STRING - []string
type deck []string

//function returns a deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	//Replace VARIABLE with UNDERSCORE if NOT USED
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//RECEIVER FUNCTION
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Return First five and remianing cards as separate variable
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//RECEIVER FUNCTION
//deck -> []string -> string -> []byte
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//RECEIVER FUNCTION
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//option 1 - log error -> call newDeck()
		//option 2 - log error and quit
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",") //convert byte slice into a string
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UTC().UnixNano())
	//Create a new RAND object with random seed value
	r := rand.New(source)

	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		//swap positions of cards in deck
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
