package main

import (
	"Fusu/router"
	"Fusu/ws"
)

func main() {
	go ws.Manager()
	router.Start()
}

