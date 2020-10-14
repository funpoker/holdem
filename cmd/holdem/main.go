package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/funpoker/holdem/game"
)

func main() {
	addr := flag.String("addr", ":8080", "http service address")
	flag.Parse()

	roomManager := game.NewRoomManager([]int{1001, 1002})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		game.ServeWs(roomManager, w, r)
	})
	http.Handle("/proto/", http.StripPrefix("/proto/", http.FileServer(http.Dir("proto/"))))

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
