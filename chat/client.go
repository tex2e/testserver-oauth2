
package main

import (
    "time"

    "github.com/gorilla/websocket"
)

type client struct {
    socket *websocket.Conn
    // メッセージが送られるチャンネル
    send chan *message
    // このクライアントが参加しているチャットルーム
    room *room
    // ユーザに関する情報
    userData map[string]interface{}
}

func (c *client) read() {
    for {
        var msg *message
        if err := c.socket.ReadJSON(&msg); err == nil {
            msg.When = time.Now()
            msg.Name = c.userData["name"].(string)
            c.room.forward <- msg
        } else {
            break
        }
    }
    c.socket.Close()
}

func (c *client) write() {
    for msg := range c.send {
        err := c.socket.WriteJSON(msg)
        if err != nil {
            break
        }
    }
    c.socket.Close()
}
