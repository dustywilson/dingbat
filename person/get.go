package person

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/dustywilson/dingbat/client"
	"github.com/rsms/gotalk"
)

func init() {
	gotalk.Handle("person.get", func(s *gotalk.Sock, r *GetRequest) (*Person, error) { fmt.Println(r.ID); return get(s, *r.ID) })
}

// GetRequest is a request for a Person
type GetRequest struct {
	ID *bson.ObjectId
}

func get(s *gotalk.Sock, id bson.ObjectId) (*Person, error) {
	c := client.SockClient(s)
	fmt.Printf("GETREQUEST: [%s] [%+v] [%+v]\n", c.GoogleEmail, c, id)
	p := new(Person)
	err := coll.FindId(id).One(&p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
