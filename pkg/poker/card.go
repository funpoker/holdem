package poker

import "strings"

var (
	ranks = "23456789TJQKA"
	suits = "shdc"

	suitPrint = map[string]string{
		"s": "\u2660",
		"h": "\u2764",
		"d": "\u2666",
		"c": "\u2663",
	}
)

type Card string

func (c Card) String() string {
	if len(c) != 2 {
		return string(c)
	}
	return string(c[0]) + suitPrint[string(c[1])]
}

type CardList []Card

func (c CardList) String() string {
	cards := []string{}
	for _, card := range c {
		cards = append(cards, card.String())
	}

	return strings.Join(cards, " ")
}
