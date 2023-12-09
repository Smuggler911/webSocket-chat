package webSocketConf

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"webSocket-chat/types"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//функция для вывода сообщения если на дом.странице

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, we glad to see you in chat room ")
}

// обработка соединения + чтение сообщения
func HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	types.Clients[conn] = true

	for {
		var msq types.Message
		err := conn.ReadJSON(&msq)
		if err != nil {
			fmt.Println(err)
			delete(types.Clients, conn)
			return
		}
		types.BroadCast <- msq
	}
}

// / обработка сообщения
func HandleMessage() {
	for {
		msg := <-types.BroadCast

		for client := range types.Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println(err)
				client.Close()
				delete(types.Clients, client)
			}
		}
	}
}
