package conn

import (
	"log"
	"strings"
	"golang.org/x/net/websocket"
	"webtailf/file"
)

type connection struct {
	ws	*websocket.Conn
	send 	chan string
}

func GetConnection(ws *websocket.Conn) *connection{
	return &connection {send: make(chan string, 8182), ws: ws}
}

func (this *connection) Reader() {
	for {
		var message string
		err := websocket.Message.Receive(this.ws, &message)
		if err != nil {
			log.Print("Receive-Error:", err)
			break
		}
	}
	this.ws.Close()
}

func (this *connection) Writer() {

	// write last 10 lines and then wait for any update
        lines, _ := file.ReadLastNLines(10)
        content := strings.Join(lines, "\n")
        this.send <- content

	for message := range this.send {
		err := websocket.Message.Send(this.ws, message)
		if err != nil {
			log.Print("Send-Error:", err)
			break
		}
	}
	this.ws.Close()
}
