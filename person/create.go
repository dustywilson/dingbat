package person

import (
	"fmt"

	"github.com/dustywilson/dingbat/db"
	"github.com/rsms/gotalk"
)

func init() {
	gotalk.Handle("person.create", create)
}

func create(s *gotalk.Sock, r *Person) (*Person, error) {
	fmt.Printf("CREATEREQUEST: [%+v] [%+v]\n", s.UserData, r)

	r.ID = db.NewObjectID()

	// FIXME: do other checks of this input...

	coll.Insert(r)

	return get(s, r.ID)
}