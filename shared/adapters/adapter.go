package adapters

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// DefaultAdapters creates a list containing the different adaptors available
var DefaultAdapters = []string{"shell", "http"}

// // Refactor 2
// // NOTE: A map usually doesn't have any values. To initialize one, we use golang built-in make command
// Grades := make(map[string]float32)

// Grades["Timmy"] = 42
// Grades["Jess"] = 92
// Grades["Sam"] = 67
// fmt.Println(Grades)

// Adapter describes an interface all Adapters should implement
// TODO: Re-enable
// type Adapter interface {
// 	// ValidateRequest(r *http.Request) (validated bool, err error)

// }

// EXAMPLE:
// // ValidateRequest validates that a github user belongs to a team on the allowedTeams list
// func (ga *GithubAuthorizer) ValidateRequest(r *http.Request) (validated bool, err error) {
// 	tokenVal, err := ParseToken(r)
// 	if err != nil {
// 		return false, err
// 	}

// 	token := oauth2.Token{AccessToken: tokenVal}
// 	return ga.validateUser(&token)
// }

// {EventEmitter} = require 'events'

// class Adapter extends EventEmitter
//   # An adapter is a specific interface to a chat source for robots.
//   #
//   # robot - A Robot instance.
//   constructor: (@robot) ->

//   # Public: Raw method for sending data back to the chat source. Extend this.
//   #
//   # envelope - A Object with message, room and user details.
//   # strings  - One or more Strings for each message to send.
//   #
//   # Returns nothing.
//   send: (envelope, strings...) ->

//   # Public: Raw method for sending emote data back to the chat source.
//   # Defaults as an alias for send
//   #
//   # envelope - A Object with message, room and user details.
//   # strings  - One or more Strings for each message to send.
//   #
//   # Returns nothing.
//   emote: (envelope, strings...) ->
//     @send envelope, strings...

//   # Public: Raw method for building a reply and sending it back to the chat
//   # source. Extend this.
//   #
//   # envelope - A Object with message, room and user details.
//   # strings  - One or more Strings for each reply to send.
//   #
//   # Returns nothing.
//   reply: (envelope, strings...) ->

//   # Public: Raw method for setting a topic on the chat source. Extend this.
//   #
//   # envelope - A Object with message, room and user details.
//   # strings  - One more more Strings to set as the topic.
//   #
//   # Returns nothing.
//   topic: (envelope, strings...) ->

//   # Public: Raw method for playing a sound in the chat source. Extend this.
//   #
//   # envelope - A Object with message, room and user details.
//   # strings  - One or more strings for each play message to send.
//   #
//   # Returns nothing
//   play: (envelope, strings...) ->

//   # Public: Raw method for invoking the bot to run. Extend this.
//   #
//   # Returns nothing.
//   run: ->

//   # Public: Raw method for shutting the bot down. Extend this.
//   #
//   # Returns nothing.
//   close: ->

//   # Public: Dispatch a received message to the robot.
//   #
//   # Returns nothing.
//   receive: (message) ->
//     @robot.receive message

//   # Public: Get an Array of User objects stored in the brain.
//   #
//   # Returns an Array of User objects.
//   users: ->
//     @robot.logger.warning '@users() is going to be deprecated in 3.0.0 use @robot.brain.users()'
//     @robot.brain.users()

//   # Public: Get a User object given a unique identifier.
//   #
//   # Returns a User instance of the specified user.
//   userForId: (id, options) ->
//     @robot.logger.warning '@userForId() is going to be deprecated in 3.0.0 use @robot.brain.userForId()'
//     @robot.brain.userForId id, options

//   # Public: Get a User object given a name.
//   #
//   # Returns a User instance for the user with the specified name.
//   userForName: (name) ->
//     @robot.logger.warning '@userForName() is going to be deprecated in 3.0.0 use @robot.brain.userForName()'
//     @robot.brain.userForName name

//   # Public: Get all users whose names match fuzzyName. Currently, match
//   # means 'starts with', but this could be extended to match initials,
//   # nicknames, etc.
//   #
//   # Returns an Array of User instances matching the fuzzy name.
//   usersForRawFuzzyName: (fuzzyName) ->
//     @robot.logger.warning '@userForRawFuzzyName() is going to be deprecated in 3.0.0 use @robot.brain.userForRawFuzzyName()'
//     @robot.brain.usersForRawFuzzyName fuzzyName

//   # Public: If fuzzyName is an exact match for a user, returns an array with
//   # just that user. Otherwise, returns an array of all users for which
//   # fuzzyName is a raw fuzzy match (see usersForRawFuzzyName).
//   #
//   # Returns an Array of User instances matching the fuzzy name.
//   usersForFuzzyName: (fuzzyName) ->
//     @robot.logger.warning '@userForFuzzyName() is going to be deprecated in 3.0.0 use @robot.brain.userForFuzzyName()'
//     @robot.brain.usersForFuzzyName fuzzyName

//   # Public: Creates a scoped http client with chainable methods for
//   # modifying the request. This doesn't actually make a request though.
//   # Once your request is assembled, you can call `get()`/`post()`/etc to
//   # send the request.
//   #
//   # Returns a ScopedClient instance.
//   http: (url) ->
//     @robot.logger.warning '@http() is going to be deprecated in 3.0.0 use @robot.http()'
//     @robot.http(url)

// module.exports = Adapter
