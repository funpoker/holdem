package game

import "sync"

type RoomManager struct {
	sync.RWMutex

	rooms map[int]*RoomHub
}

func NewRoomManager(roomIds []int) *RoomManager {
	rm := &RoomManager{
		rooms: make(map[int]*RoomHub),
	}

	for _, id := range roomIds {
		rm.Add(id)
	}

	return rm
}

func (r *RoomManager) RoomHub(roomId int) *RoomHub {
	r.RLock()
	defer r.RUnlock()

	return r.rooms[roomId]
}

func (r *RoomManager) Add(roomId int) {
	r.Lock()
	defer r.Unlock()

	r.rooms[roomId] = NewRoomHub()
}
