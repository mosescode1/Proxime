package tile38

import (
	"github.com/xjem/t38c"
	"log"
)

var T38Client *t38c.Client

func NewT38Client() (*t38c.Client, error) {
	t38, err := t38c.New(t38c.Config{
		Address: "localhost:9851",
		Debug:   true,
	})
	if err != nil {
		return nil, err
	}

	T38Client = t38
	return t38, nil
}

func T38Close(client *t38c.Client) {
	err := client.Close()
	if err != nil {
		log.Fatal(err)
	}
}
