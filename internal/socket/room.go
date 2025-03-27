package socket

type Room struct {
	name    string
	clients map[*Client]bool
}
