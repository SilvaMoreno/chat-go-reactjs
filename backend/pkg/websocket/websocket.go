package websocket

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// we'll need to define an Upgrader
// this will requere a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// we'll need to check the origin of our connection
	// this will allow us to make request from our frontend
	CheckOrigin: func(r *http.Request) bool { return true },
}

// define websocket endpoint
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	// fmt.Println(r.Host)

	// upgrade this connection to a WebSocket
	// connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//listen indefinitely fron new messages coming
	// through on our WebSocket connection
	return conn, nil
}

// define a reader which will listen for new messages sended to webSocket
func Reader(conn *websocket.Conn) {
	for {
		//read in a message
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		// print message
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func Writer(conn *websocket.Conn) {
	for {
		//set in a message
		fmt.Println("Sending")
		messageType, r, err := conn.NextReader()

		if err != nil {
			log.Println(err)
			return
		}

		w, err := conn.NextWriter(messageType)

		if err != nil {
			log.Println(err)
			return
		}

		if _, err := io.Copy(w, r); err != nil {
			fmt.Println(err)
			return
		}

		if err := w.Close(); err != nil {
			log.Println(err)
			return
		}
	}
}
