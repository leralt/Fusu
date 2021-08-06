package main

import (
	"Fusu/ws"
	"net/http"
)

func main() {
	go ws.Manager()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/room.html")
	})

	http.HandleFunc("/room", ws.EchoMessage)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		return
	}
}
