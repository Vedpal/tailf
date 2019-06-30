package conn

import (
    	"sync"
)

type connection_manager struct {
	connections	map[*connection]bool
	register 	chan *connection
	unregister 	chan *connection
	broadcast 	chan string
}

var once sync.Once
var connMgr *connection_manager

func GetConnMgr() *connection_manager {
    once.Do(func() {
        connMgr = &connection_manager{
        	connections: make(map[*connection]bool),
        	broadcast:   make(chan string),
        	register:    make(chan *connection),
        	unregister:  make(chan *connection),
	}
    })
    return connMgr
}

func (this *connection_manager) Register(c *connection) {
	connMgr.register <- c
}

func (this *connection_manager) Unregister(c *connection) {
	connMgr.unregister <- c
}

func (this *connection_manager) Broadcast(content string) {
	connMgr.broadcast <- content
}

func (this *connection_manager) Run() {
	for {
		select {
		case c := <-this.register:
			this.connections[c] = true
		case c := <-this.unregister:
			delete(this.connections, c)
			close(c.send)
		case m := <-this.broadcast:
			for c := range this.connections {
				select {
				case c.send <- m:
				default:
					delete(this.connections, c)
					close(c.send)
					go c.ws.Close()
				}
			}
		}
	}
}
