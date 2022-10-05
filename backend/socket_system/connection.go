package socket_system

import (
	"fmt"
	"log"
	"net/http"
	"time"

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
type connectionData struct {
	room string
	name string
}

// var connectionList []connectionMessage
var connectionMap map[*websocket.Conn]connectionData

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
		conn.SetReadDeadline(time.Now().Add(1 * time.Minute))
		typeMes, p, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			delete(connectionMap, conn)
			fmt.Println("reader", err)
			return
		}
		cm := ChatMessage{
			message: string(p),
			room:    room,
			from:    user,
			typeMes: typeMes,
		}
		// _ := cm
		fmt.Println("typeMes", cm)
		// if err := conn.WriteMessage(typeMes, p); err != nil {

		// 	conn.Close()
		// 	// delete(connectionMap, key)
		// }
		broadcasterChat <- cm

	}
}

type ChatMessage struct {
	message string
	room    string
	from    string
	typeMes int
}

var broadcasterChat = make(chan ChatMessage)

func handleMessage() {
	for {
		msg := <-broadcasterChat
		fmt.Println(msg)
		// for key, connect := range connectionList {
		// 	if connect.room == msg.room && connect.name != msg.from {
		// 		if err := connect.websocket.WriteMessage(0, []byte(msg.message)); err != nil {
		// 			connect.websocket.Close()
		// 			remove(connectionList, i)
		// 		}
		// 	}
		// }
		for key, val := range connectionMap {
			if val.room == msg.room && val.name != msg.from {
				if err := key.WriteMessage(msg.typeMes, []byte(msg.message)); err != nil {
					fmt.Println("errHandleMess", err)
					key.Close()
					delete(connectionMap, key)
				}
			}
		}

	}
}

func remove(slice []connectionMessage, s int) []connectionMessage {
	return append(slice[:s], slice[s+1:]...)
}
