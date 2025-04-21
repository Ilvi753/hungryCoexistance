package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Create a deck of cards
func createDeck() []string {
	suits := []string{"♠", "♥", "♦", "♣"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	var deck []string
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, rank+suit)
		}
	}
	return deck
}

// Fisher-Yates shuffle using crypto/rand for secure randomness
// Double Shuffle (shuffle -> shuffle again) to complicate the process
func shuffleDeck(deck []string) error {
	// First shuffle pass
	for i := len(deck) - 1; i > 0; i-- {
		// Generate a random index j in the range [0, i]
		jBig, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return err
		}
		j := int(jBig.Int64())

		// Swap the elements at index i and j
		deck[i], deck[j] = deck[j], deck[i]
	}

	// Second shuffle pass (to add another layer of mixing)
	for i := len(deck) - 1; i > 0; i-- {
		// Generate a random index j in the range [0, i]
		jBig, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return err
		}
		j := int(jBig.Int64())

		// Swap the elements at index i and j
		deck[i], deck[j] = deck[j], deck[i]
	}

	return nil
}

func main() {
	// Create a deck of cards
	deck := createDeck()

	// Shuffle the deck with double encryption for added complexity
	err := shuffleDeck(deck)
	if err != nil {
		fmt.Println("Error shuffling deck:", err)
		return
	}

	// Print the shuffled deck (or use it for your poker game)
	fmt.Println("Shuffled deck of cards:")
	for _, card := range deck {
		fmt.Println(card)
	}
}
