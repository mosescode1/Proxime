package socket

import (
	"errors"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	location string
	send     chan []byte
}

func NewClient(conn *websocket.Conn, location string) (*Client, error) {
	if conn == nil || location == "" {
		return &Client{}, errors.New("connection is nil or location is empty")
	}
	return &Client{
		conn:     conn,
		location: location,
		send:     make(chan []byte),
	}, nil
}
