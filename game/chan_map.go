package game

import "sync"

var chMap *ChanMap

func init() {
	chMap = NewChanMap()
}

type ChanMap struct {
	sync.RWMutex
	chs map[int]chan []byte
}

func NewChanMap() *ChanMap {
	return &ChanMap{
		chs: make(map[int]chan []byte),
	}
}

func (c *ChanMap) Add(playerId int, ch chan []byte) {
	c.Lock()
	defer c.Unlock()
	c.chs[playerId] = ch
}

func (c *ChanMap) Get(playerId int) (chan []byte, bool) {
	c.RLock()
	defer c.RUnlock()
	ch, ok := c.chs[playerId]
	return ch, ok
}
