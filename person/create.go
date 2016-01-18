package person

import (
	"fmt"

	"github.com/dustywilson/dingbat/client"
	"github.com/dustywilson/dingbat/db"
	"github.com/rsms/gotalk"
)

func init() {
	gotalk.Handle("person.create", create)
}

func create(s *gotalk.Sock, r *Person) (*Person, error) {
	c := client.SockClient(s)
	fmt.Printf("CREATEREQUEST: [%s] [%+v] [%+v]\n", c.GoogleEmail, c, r)

	r.ID = db.NewObjectID()

	// FIXME: do other checks of this input...

	coll.Insert(r)

	return get(s, r.ID)
}
