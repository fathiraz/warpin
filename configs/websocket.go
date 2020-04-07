package configs

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	// websocket var
	ws []*WebsocketConnections
)

type WebsocketConnections struct {
	*websocket.Conn
	Name string
}

func AddConnections(ec echo.Context, name string) error {
	var upgrader = websocket.Upgrader{}

	// set new websocket with upgrade our http
	newWs, err := upgrader.Upgrade(ec.Response(), ec.Request(), nil)
	if err != nil {
		return err
	}

	// set new connection from our currenlty ws connection
	newConnection := WebsocketConnections{
		Conn: newWs,
		Name: name,
	}

	// append to our pool of ws connections
	ws = append(ws, &newConnection)

	return nil
}

func GetConnectionsByName(name string) *WebsocketConnections {
	for _, wsConnections := range ws {
		if name == wsConnections.Name {
			return wsConnections
		}
	}

	return nil
}

func Broadcast(name string, data interface{}) {
	conn := GetConnectionsByName(name)

	// check for connection not null
	if conn != nil {

		// loop to broadcast to all connection
		for _, webs := range ws {

			// if connection our connection, no need to broadcast
			if conn.Name != webs.Name {
				// set message to broadcast with chat socket response
				webs.Conn.WriteJSON(data)
			}
		}
	}
}
