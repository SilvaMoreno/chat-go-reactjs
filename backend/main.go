package main

import (
	"fmt"
	"net/http"

	"github.com/SilvaMoreno/realtime-chat-go-reactjs/pkg/websocket"
)

// define websocket endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	fmt.Println(r.Host)

	// upgrade this connection to a WebSocket
	// connection
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Println(w, "%+V\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
