package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/rsms/gotalk"
	"github.com/satori/go.uuid"
)

// Client is a connected web client
type Client struct {
	id       uuid.UUID
	lastSeen time.Time
}

var clients = struct {
	clients map[string]*Client
	sync.RWMutex
}{
	clients: make(map[string]*Client),
}

func newClient() *Client {
	clients.Lock()
	defer clients.Unlock()

	c := &Client{
		id:       uuid.NewV4(),
		lastSeen: time.Now(),
	}

	clients.clients[c.id.String()] = c

	return c
}

func getClient(id uuid.UUID) *Client {
	return clients.clients[id.String()]
}

func onAccept(s *gotalk.Sock) {
	c := newClient()
	s.UserData = c
	s.Notify("connection.id", c.id.String())
	fmt.Printf("New accepted connection.  [%s]\n", c.id)
}

//struct{ ID string }{ID:

// Touch a client to update its lastSeen value
func (c *Client) Touch() {
	c.lastSeen = time.Now()
}
