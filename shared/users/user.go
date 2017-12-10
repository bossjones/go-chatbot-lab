package users

import (
	"fmt"

	util "github.com/bossjones/go-chatbot-lab/shared/lib"
)

// User represents a participating user in the chat.
type User struct {
	// Id - A unique ID for the user.
	id string

	// Options - An optional Hash of key, value pairs for this user.
	options map[string]string
}

// NewUser returns a reference to an instance of User
func NewUser(op map[string]string) *User {

	user := User{
		id:      util.CreateUUID(),
		options: make(map[string]string),
	}

	fmt.Printf("id %s", user.id)

	return &user
}
