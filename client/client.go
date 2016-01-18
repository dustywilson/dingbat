package client

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"google.golang.org/api/oauth2/v1"

	"github.com/rsms/gotalk"
	"github.com/satori/go.uuid"
)

var googleClientID = ""
var oa2 *oauth2.Service

// Client is a connected web client
type Client struct {
	valid               bool
	ID                  uuid.UUID
	LastSeen            time.Time
	googleToken         *oauth2.Tokeninfo
	GoogleUserID        string
	GoogleEmail         string
	GoogleEmailVerified bool
}

var clients = struct {
	clients map[string]*Client
	sync.RWMutex
}{
	clients: make(map[string]*Client),
}

func init() {
	gotalk.Handle("client.init-google", sessionLoginGoogle)
	gotalk.Handle("client.signout-google", sessionLogoutGoogle)

	if envGoogleClientID := os.Getenv("GOOGLE_CLIENTID"); envGoogleClientID != "" {
		googleClientID = envGoogleClientID
	}

	if len(googleClientID) == 0 {
		panic("GOOGLE_CLIENTID must not be empty.")
	}

	var err error
	oa2, err = oauth2.New(&http.Client{})
	if err != nil {
		panic(err)
	}
}

// NewClient returns a new *Client with a fresh ID
func NewClient() *Client {
	clients.Lock()
	defer clients.Unlock()

	c := &Client{
		valid:    true,
		ID:       uuid.NewV4(),
		LastSeen: time.Now(),
	}

	clients.clients[c.ID.String()] = c

	return c
}

// SockClient extracts the *Client from the socket state
func SockClient(s *gotalk.Sock) *Client {
	if s == nil || s.UserData == nil {
		return nil
	}
	return s.UserData.(*Client)
}

func getClient(id uuid.UUID) *Client {
	return clients.clients[id.String()]
}

func sessionLoginGoogle(s *gotalk.Sock, googleToken string) (*Client, error) {
	c := SockClient(s)
	fmt.Printf("Google Login: [%+v] [%+v] [%+v]\n", googleClientID, googleToken, c)
	t, err := oa2.Tokeninfo().IdToken(googleToken).Do()
	if err != nil {
		return nil, err
	}
	fmt.Printf("TOKENINFO:[%+v]\n", t)
	if t.Audience != googleClientID {
		return nil, errors.New("Token error")
	}
	if t.Issuer != "accounts.google.com" && t.Issuer != "https://accounts.google.com" {
		return nil, errors.New("Token error")
	}
	c.GoogleUserID = t.UserId
	c.GoogleEmail = t.Email
	c.GoogleEmailVerified = t.EmailVerified
	c.googleToken = t
	return c, nil
}

func sessionLogoutGoogle(s *gotalk.Sock, _ interface{}) error {
	c := SockClient(s)
	c.valid = false
	fmt.Printf("Google Logout: [%+v]\n", c)
	// TODO: should remove the session from the store
	s.Close()
	return nil
}

// Touch a client to update its LastSeen value
func (c *Client) Touch() {
	c.LastSeen = time.Now()
}

// IsValid returns whether or not the Client is valid (not expired, invalidated)
func (c *Client) IsValid() bool {
	return c.valid
}
