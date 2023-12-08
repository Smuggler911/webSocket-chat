package main

import (
	"fmt"
	"net/http"
	"webSocket-chat/webSocketConf"
)

func main() {
	http.HandleFunc("/", webSocketConf.HomePage)
	http.HandleFunc("/ws", webSocketConf.HandleConnection)

	go webSocketConf.HandleMessage()

	fmt.Println("Server started on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
