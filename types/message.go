package types

import "github.com/gorilla/websocket"

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var Clients = make(map[*websocket.Conn]bool)
var BroadCast = make(chan Message)
