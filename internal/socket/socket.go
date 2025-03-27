package socket

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"log"
	"log/slog"
	"net/http"
)

var wsUpgrade = &websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins; adjust as needed for security
	},
}

type Socket struct {
	rooms map[string]*Room
}

func NewSocket() *Socket {
	return &Socket{
		rooms: make(map[string]*Room),
	}
}

func (s *Socket) GetRooms() map[string]*Room {

	slog.Info("GetRooms", slog.Int("rooms", len(s.rooms)))
	return s.rooms
}

func (s *Socket) Connect(c echo.Context) error {
	conn, err := wsUpgrade.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	defer func(ws *websocket.Conn) {
		if err := ws.Close(); err != nil {
			slog.Error(err.Error(), slog.String("socket", ws.RemoteAddr().String()))
		}
	}(conn)

	location := c.Request().URL.Query().Get("room")

	// create a client
	client, err := NewClient(conn, location)

	if err != nil {
		slog.Error(err.Error(), slog.String("Cannot Create Client ", conn.RemoteAddr().String()))
		return err
	}

	s.addClientToRoom(client)

	// Start reading messages from client
	go s.handleMessage(client)

	// Write messages to the client
	for msg := range client.send {
		err := client.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}

	s.removeClientFromRoom(client)
	return nil
}

func (s *Socket) addClientToRoom(client *Client) {
	room, exists := s.rooms[client.location]
	if !exists {
		room = &Room{
			name:    client.location,
			clients: make(map[*Client]bool),
		}
		s.rooms[client.location] = room
	}
	room.clients[client] = true
}

func (s *Socket) removeClientFromRoom(client *Client) {
	room, exists := s.rooms[client.location]
	if exists {
		delete(room.clients, client)
		if len(room.clients) == 0 {
			delete(s.rooms, client.location)
		}
	}
}

func (s *Socket) handleMessage(client *Client) {
	for {
		messageType, msg, err := client.conn.ReadMessage()

		if websocket.IsCloseError(err) {
			log.Println("websocket closed by client", client.conn.RemoteAddr().String())
			break
		}

		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		if messageType == websocket.TextMessage {
			s.broadcastMessage(client.location, msg, client)
		}
	}
}

func (s *Socket) broadcastMessage(roomName string, message []byte, sender *Client) {

	room, exists := s.rooms[roomName]
	if !exists {
		return
	}

	for client := range room.clients {
		// Skip sending the message to the sender
		if client == sender {
			continue
		}
		select {
		case client.send <- message:
		default:
			delete(room.clients, client)
		}
	}
}
