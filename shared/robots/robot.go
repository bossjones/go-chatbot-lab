package robots

// AliasSentinel sentinel to use by default
const AliasSentinel string = "ALIAS_SENT"

// Robot represents a chatbot Robot based on Hubot Interface (adapterPath, adapter, httpd, name = 'Hubot', alias = false)
type Robot struct {
	// adapterPath is a String of the path to built-in adapters (defaults to src/adapters)
	adapterPath string

	// adapter is a String of the adapter name.
	adapter string

	// httpd is a Boolean whether to enable the HTTP daemon
	httpd bool

	// name is a String of the robot name, defaults to Hubot.
	name string

	// alias is an alternative name to use for your chatbot
	alias string

	// Brain
}

// type car struct {
// 	gas_pedal uint16 // anything from 0 - 65535
// 	break_pedal uint16
// 	steering_wheel int16 // -32K - +32k
// 	top_speed_kmh float64
// }

// instantiate a new GithubAuthorizer
// func newRobot(config config.Config) (*Robot, error) {
// 	githubAPIURL, err := parseGithubAPI(config.GithubAPI)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ga := &GithubAuthorizer{
// 		allowedTeams: config.GithubAllowedTeams,
// 		githubAPI:    githubAPIURL,
// 	}

// 	return ga, nil
// }
