package main
import (
	"log"
	"golang.org/x/net/websocket"
	"tailf/conn"
)

func wsHandler(ws *websocket.Conn) {
	log.Print("Handling-Request...")
        c := conn.GetConnection(ws);

	log.Print("Websocket-Client-Address: ", ws.RemoteAddr())
        conn.GetConnMgr().Register(c)

        defer func() { conn.GetConnMgr().Unregister(c) }()

        go c.Writer()
        c.Reader()
}
