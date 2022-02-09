package server

import (
	"sync"
	"time"

	"math/rand"

	"github.com/gorilla/websocket"
	// "encoding/json"
	// "log"
)

type Participant struct {
	Host bool
	Conn *websocket.Conn
}

type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participant)
}

func (r *RoomMap) Get(roomID string) []Participant {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.Map[roomID]
}

func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

}

func (r *RoomMap) DeleteRoom() {

}
