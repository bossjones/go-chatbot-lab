package brains

import (
	"github.com/bossjones/go-chatbot-lab/shared/robots"
	"errors"
	"strconv"
	"strings"
	"sync"
)

// ***************************************************************************
// var (
// 	ErrKeyNotFound     = errors.New("Key not found")
// 	ErrCouldNotConnect = errors.New("Could not connect to KV store")
// )

// KVWrapper is an interface that any Key Value Store (etcd, consul) needs to implement
// when used by flight director.
// type KVWrapper interface {
// 	NewKVWrapper(servers []string, username, password string) KVWrapper
// 	Set(key string, val string, ttl uint64) error
// 	GetVal(key string) (*KeyValue, error)
// 	GetList(key string, sort bool) ([]*KeyValue, error)
// }

// func NewKVWrapper(servers []string, wrapper KVWrapper) KVWrapper {
// 	kvw := wrapper.NewKVWrapper(servers, "", "")
// 	return kvw
// }
// // ***************************************************************************

// // KeyValue entity represents the unit returned by queries to a Key Value store.
// type KeyValue struct {
// 	Key         string
// 	Value       string
// 	HasChildren bool
// }

// func (kv *KeyValue) String() string {
// 	return kv.Key + " : " + kv.Value + " : " + strconv.FormatBool(kv.HasChildren)
// }

// Data is an object containing properties data.user = {} and data.private = {}
type Data struct {
	user    map[string]string
	private map[string]string
}

// // DataValue entity represents the unit returned by queries to a Key Value store.
// type DataValue struct {
// 	Key         string
// 	Value       string
// 	HasChildren bool
// }

// func (dv *DataValue) String() string {
// 	return dv.Key + " : " + dv.Value + " : " + strconv.FormatBool(dv.HasChildren)
// }

// Brain represents a somewhat persistent storage for the robot. Extend this. ( Returns a new Brain with no external storage. )
type Brain struct {
	data *Data

	// EXAMPLE:
	// data =
	//   users:    { }
	//   _private: { }

	autoSave bool

	// EXAMPLE:
	// robot.on "running", =>
	//   @resetSaveInterval 5
}

// instantiate a new GithubAuthorizer
func newBrain(robot robots.Robot) (*Brain, error) {
	// TODO: We need to pass in some sort of struct that represents commandline data
	// githubAPIURL, err := parseGithubAPI(config.GithubAPI)
	// if err != nil {
	// 	return nil, err
	// }

	brain := &Brain{
		data:     Data,
		autoSave: false,
	}

	return brain, nil
}

// Value Receiver Method
func (b *brain) set(newvale string){
	// b.

	// Grades := make(map[string]float32)

	// Grades["Timmy"] = 42
	// NOTE: If you don't make this into a pointer receiver, then setting value below will not persist other calls!
	c.top_speed_kmh = 500
	// Returns top speed of car in km/h
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbixmax)
}

// // Reading through pointer to get value, then modify it
// func (c *car) new_top_speed(newspeed float64){
// 	c.top_speed_kmh = newspeed
// }

// // Value Receiver Method
// func (b *Brain) set() float64 {
// 	// Returns top speed of car in km/h
// 	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbixmax)
// }

// String() is a custom method that returns the Config with sensitive values
// filtered out
// func (c Config) String() string {
// 	sanitize := true
// 	return string(c.JSON(sanitize))
// }
