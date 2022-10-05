package socket_system

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	go handleMessage()
	return &Handler{service}
}
func (h *Handler) ConnectRoom(c *gin.Context) {
	var ws *websocket.Conn
	room := c.Param("room")
	if connectionMap == nil {
		connectionMap = make(map[*websocket.Conn]connectionData)
	}
	userId := c.Param("id")
	// if room == "" {
	// 	return
	// }
	// for _, connection := range connectionList {
	// 	if connection.name == room {
	// 		ws = connection.websocket
	// 	}
	// }
	if ws == nil {
		socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		connectionMap[socket] = connectionData{room: room, name: userId}
		// connectionList = append(connectionList, connectionType{websocket: socket, name: room})
		ws = socket
	}

	Reader(ws, room, userId)
}
