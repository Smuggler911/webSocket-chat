package types

import "github.com/gorilla/websocket"

// тип послыаемого сообщения
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

// клиенты
var Clients = make(map[*websocket.Conn]bool)

// канал для вывода сообщения
var BroadCast = make(chan Message)
