package types

import "github.com/gorilla/websocket"

// типа послыаемого сообщения
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

// клиенты
var Clients = make(map[*websocket.Conn]bool)

// канал для вывод сообщения
var BroadCast = make(chan Message)
