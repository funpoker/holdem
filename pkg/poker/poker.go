package poker

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

var cards = []Card{}

func init() {
	for _, s := range strings.Split(suits, "") {
		for _, r := range strings.Split(ranks, "") {
			cards = append(cards, Card(r+s))
		}
	}
}

type Poker struct {
	sync.RWMutex
	Idx   int
	Cards []Card
}

func New() *Poker {
	p := &Poker{
		Cards: cards,
	}

	return p
}

func (p *Poker) Shuffle() {
	p.Lock()
	defer p.Unlock()

	randCards := make([]Card, len(p.Cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(p.Cards))
	for i, randIndex := range perm {
		randCards[i] = p.Cards[randIndex]
	}
	p.Cards = randCards
}

func (p *Poker) Get() Card {
	p.Lock()
	defer p.Unlock()

	if p.Idx >= len(p.Cards) {
		return ""
	}

	c := p.Cards[p.Idx]
	p.Idx += 1
	return c
}
