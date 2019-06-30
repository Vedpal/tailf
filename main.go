package main

import (
	"flag"
	"log"
	"net/http"
	"golang.org/x/net/websocket"
	"tailf/conn"
	"tailf/tail"
	"tailf/file"
)

var port = flag.String("port", ":8080", "HTTP port to listen at")
var filename = flag.String("file", "/var/log/system.log", "the log filename to tail -f")

func main() {

	// parse command line options
	flag.Parse()

	// set file name
	file.SetFileName(*filename)

	// connection manager
	go conn.GetConnMgr().Run()

	// watch and broadcast tail latest line to connection manager
	go tail.Filetail()

	// web socket handler registeration
	http.Handle("/tailf", websocket.Handler(wsHandler))
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
