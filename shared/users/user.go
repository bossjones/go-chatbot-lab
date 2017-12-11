//go:generate mockgen -destination mock/user_mock.go github.com/bossjones/go-chatbot-lab/shared/users User

// source: https://blog.carlmjohnson.net/post/2016-11-27-how-to-use-go-generate/

// Please consult the docs for a full specification of the flags and options it accepts. When you run the command go generate, the Go tool scans the files relevant to the current package for lines with a “magic comment” of the form //go:generate command arguments. This command does not have to do anything related to Go or code generation.

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
