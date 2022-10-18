package route

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)


type Message struct {
	Message string `json:"msg"`
	ClientId string `json:"client_id"`
	RoomId string `json:"room_id"`
}

type WebSocket struct {
	Upgrader *websocket.Upgrader
	Write int
	Read int
	Conn *websocket.Conn
	err error

}


func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	var ws WebSocket

	ws.Upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws.Upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	ws.Conn, ws.err = ws.Upgrader.Upgrade(w, r, nil)

	if ws.err != nil {
		log.Fatalf("%v",ws.err)
	}
	
	defer ws.Conn.Close()

	for {
		var Message Message
		err := ws.Conn.ReadJSON(&Message)
		if err != nil {
			log.Fatalf("%v",err)
			break
		}


		fmt.Printf("Message Received: %v\n",Message.Message)
	}

}

func NewRouter()  {
	router := mux.NewRouter()

	router.HandleFunc("/socket", WsEndpoint)
	log.Fatal(http.ListenAndServe(":8080", router))
}