package socket_system

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type connectionMessage struct {
	websocket *websocket.Conn
	room      string
	name      string
}

var connectionList []connectionMessage

// type connection struct {
// }

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	defer ws.Close()
	if err != nil {
		log.Println(err)
	}
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}
	// reader(ws)

}
func Reader(conn *websocket.Conn, room string, user string) {
	for {
		// read in a message
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		cm := ChatMessage{
			message: string(p),
			room:    room,
			from:    user,
		}
		broadcasterChat <- cm

	}
}

type ChatMessage struct {
	message string
	room    string
	from    string
}

var broadcasterChat = make(chan ChatMessage)

func handleMessage() {
	for {
		msg := <-broadcasterChat
		for i, connect := range connectionList {
			if connect.room == msg.room && connect.name != msg.from {
				if err := connect.websocket.WriteMessage(0, []byte(msg.message)); err != nil {
					connect.websocket.Close()
					remove(connectionList, i)
				}
			}
		}

	}
}

func remove(slice []connectionMessage, s int) []connectionMessage {
	return append(slice[:s], slice[s+1:]...)
}
