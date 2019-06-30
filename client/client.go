package main

import (
	"log"
	"fmt"
	"golang.org/x/net/websocket"
)

func main() {

	origin := "http://localhost/"
	url := "ws://localhost:8080/tailf"

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
    		log.Fatal(err)
	}

	var msg = make([]byte, 8192)
	var n int
	for {
		if n, err = ws.Read(msg); err != nil {
    			log.Fatal(err)
		}
		fmt.Printf("%s\n", msg[:n])
	}
}
