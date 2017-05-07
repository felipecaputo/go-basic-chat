package main

import (
	"github.com/gorilla/websocket"
)

//client represents a single chat user
type client struct {
	//socket is the client socket for this user
	socket *websocket.Conn
	//send is a channel which messages are sent
	send chan []byte
	//room is the room where client is chatting in
	room *room
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}

	c.socket.Close()
}
