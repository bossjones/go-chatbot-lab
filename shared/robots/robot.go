package robots

// const Brain = require('./brain')
// const Response = require('./response')
// const Listener = require('./listener')
// const Message = require('./message')
// const Middleware = require('./middleware')

import (
	"github.com/bossjones/go-chatbot-lab/log"
	"github.com/bossjones/go-chatbot-lab/shared/brains"
)

// AliasSentinel sentinel to use by default
const AliasSentinel string = "ALIAS_SENT"

// DefaultName -
const DefaultName = string("Scarlett")

// DefaultPort -
const DefaultPort = string("5535")

// DefaultTTL -
const DefaultTTL = uint64(60)

// DefaultTimeout -
const DefaultTimeout = uint64(5)

// TTLRenew -
const TTLRenew = int64(30)

// ResolverAgeSec -
const ResolverAgeSec = float64(20)

// DefaultAdapter -
const DefaultAdapter = string("shell")

// ChatbotDefaultAdapters is a string slice ( similar to array, but not of fixed size ) of available adaptors
var ChatbotDefaultAdapters = []string{"campfire", "shell"}

// ChatbotDocumentationSections is a string slice ( similar to array, but not of fixed size ) of avaiable documentation sections
var ChatbotDocumentationSections = []string{"description", "dependencies", "configuration", "commands", "notes", "author", "authors", "examples", "tags", "urls"}

// Robot represents a chatbot Robot based on Hubot Interface (adapterPath, adapter, httpd, name = 'Hubot', alias = false)
type Robot struct {
	// AdapterPath is a String of the path to built-in adapters (defaults to src/adapters)
	AdapterPath *string

	// Adapter is a String of the adapter name.
	Adapter *string

	// Httpd is a Boolean whether to enable the HTTP daemon
	Httpd *bool

	// Name is a String of the robot name, defaults to Hubot.
	Name *string

	// Alias is an alternative name to use for your chatbot
	Alias *string

	// RobotBrain -
	RobotBrain *brains.Brain
}

// NewRobot returns a reference to an instance of Robot
func NewRobot(adapterpath *string, adapter *string, httpd *bool, name *string, alias *string) *Robot {

	log.WithFields(log.Fields{
		"EventName": "new_robot",
		"Info":      "Creating new Robot Object",
	}).Info("Creating new Robot Object")

	if *adapterpath == "" {
		*adapterpath = DefaultAdapter
	}

	if *adapter == "" {
		*adapter = DefaultAdapter
	}

	// if *httpd == true {
	// 	setupExpress()
	// }
	// else {
	// 	setupNullRouter()
	// }

	*httpd = false

	if *name == "" {
		*name = DefaultName
	}

	if *alias == "" {
		*alias = DefaultName
	}

	rdata := brains.NewData()
	rbrain := brains.NewBrain(rdata)

	robot := Robot{
		AdapterPath: adapterpath,
		Adapter:     adapter,
		Httpd:       httpd,
		Name:        name,
		Alias:       alias,
		RobotBrain:  rbrain,
	}

	return &robot
}

// Robots receive messages from a chat source (Campfire, irc, etc), and
// dispatch them to matching listeners.
//
// adapterPath -  A String of the path to built-in adapters (defaults to src/adapters)
// adapter     - A String of the adapter name.
// httpd       - A Boolean whether to enable the HTTP daemon.
// name        - A String of the robot name, defaults to Hubot.

// func NewPersist( hostName *string, stateHost *string, statePort *string, stateUser *string, statePass *string) *Persist {

// 	if *stateHost == "" {
// 		*stateHost = DefaultHost
// 	}

// 	if *statePort == "" {
// 		*statePort = DefaultPort
// 	}

// 	return &Persist{
// 		Hostname:  hostName,
// 		StateHost: stateHost,
// 		StatePort: statePort,
// 		StateUser: stateUser,
// 		StatePass: statePass,
// 	}

// }

// constructor (adapterPath, adapter, httpd, name, alias) {
//     if (name == null) {
//       name = 'Hubot'
//     }
//     if (alias == null) {
//       alias = false
//     }
//     this.adapterPath = path.join(__dirname, 'adapters')

//     this.name = name
//     this.events = new EventEmitter()
//     this.brain = new Brain(this)
//     this.alias = alias
//     this.adapter = null
//     this.Response = Response
//     this.commands = []
//     this.listeners = []
//     this.middleware = {
//       listener: new Middleware(this),
//       response: new Middleware(this),
//       receive: new Middleware(this)
//     }
//     this.logger = new Log(process.env.HUBOT_LOG_LEVEL || 'info')
//     this.pingIntervalId = null
//     this.globalHttpOptions = {}

//     this.parseVersion()
//     if (httpd) {
//       this.setupExpress()
//     } else {
//       this.setupNullRouter()
//     }

//     this.loadAdapter(adapter)

//     this.adapterName = adapter
//     this.errorHandlers = []

//     this.on('error', (err, res) => {
//       return this.invokeErrorHandlers(err, res)
//     })
//     this.onUncaughtException = err => {
//       return this.emit('error', err)
//     }
//     process.on('uncaughtException', this.onUncaughtException)
//   }
