// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "GitHub SSH Keys": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/JKhawaja/replicated/design
// --out=$(GOPATH)\src\github.com\JKhawaja\replicated
// --version=v1.2.0-dirty

package app

import (
	"encoding/json"

	"github.com/goadesign/goa"
)

// Response Type for a GitHub User's list of public SSH Keys (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	// The list of the Github user's public SSH keys.
	Keys []*UserKey `form:"keys" json:"keys" xml:"keys"`
	// The username of the GitHub user.
	Username string `form:"username" json:"username" xml:"username"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {
	if mt.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}
	if mt.Keys == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "keys"))
	}
	for _, e := range mt.Keys {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// UserCollection is the media type for an array of User (default view)
//
// Identifier: application/vnd.user+json; type=collection; view=default
type UserCollection []*User

func (u UserCollection) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{}, 0)

	for _, n := range u {
		data[n.Username] = n.Keys
	}

	return json.Marshal(data)
}

// Validate validates the UserCollection media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
