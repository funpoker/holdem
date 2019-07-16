package game

import (
	"errors"
	"fmt"

	"github.com/funpoker/holdem/pkg/poker"
)

var ErrPlayerNum = errors.New("palery number should be [2, 9]")

type HoldemPoker struct {
	poker *poker.Poker

	Flop  poker.CardList
	Turn  poker.Card
	River poker.Card

	PlayerNum int
	Hands     poker.CardList // 2*PlayerNum
}

func NewHoldemPoker(playerNum int) (*HoldemPoker, error) {
	if playerNum < 2 || playerNum > 9 {
		return nil, ErrPlayerNum
	}

	hp := &HoldemPoker{
		poker:     poker.New(),
		PlayerNum: playerNum,
	}

	hp.poker.Shuffle()

	for i := 0; i < 3; i++ {
		hp.Flop = append(hp.Flop, hp.poker.Get())
	}
	hp.Turn = hp.poker.Get()
	hp.River = hp.poker.Get()

	for i := 0; i < playerNum*2; i++ {
		hp.Hands = append(hp.Hands, hp.poker.Get())
	}

	return hp, nil
}

func (h *HoldemPoker) PlayerHands(playerIdx int) string {
	return poker.CardList(h.Hands[2*playerIdx : 2*playerIdx+2]).String()
}

func (h *HoldemPoker) String() string {
	s := fmt.Sprintf("Flop: %s\nTrun: %s\nRiver: %s\n", h.Flop, h.Turn, h.River)
	for i := 0; i < h.PlayerNum; i++ {
		s += fmt.Sprintf("Player %d: %s\n", i, h.PlayerHands(i))
	}
	return s
}
