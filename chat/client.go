
package main

import (
    "github.com/gorilla/websocket"
)

type client struct {
    socket *websocket.Conn
    // メッセージが送られるチャンネル
    send chan []byte
    // このクライアントが参加しているチャットルーム
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
    c.socket.Close()
}

func (c *client) write() {
    for msg := range c.send {
        err := c.socket.WriteMessage(websocket.TextMessage, msg)
        if err != nil {
            break
        }
    }
    c.socket.Close()
}
