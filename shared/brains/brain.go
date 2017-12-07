package brains

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
	// User is a struct field that is of type map {"":""}
	User map[string]string

	// Private is a struct field that is of type map {"":""}
	Private map[string]string
}

// NOTE: https://stackoverflow.com/questions/27553399/golang-how-to-initialize-a-map-field-within-a-struct

// NewData returns a reference to an instance of Data
func NewData() *Data {

	data := Data{
		User:    make(map[string]string),
		Private: make(map[string]string),
	}

	return &data
}

// INFO: https://stackoverflow.com/questions/27455170/error-struct-type-is-not-an-expression

// NOTE: Try using this from skopos
// Schedule - list of maintenance windows
// type Schedule struct {
// 	Windows []*Window `json:"windows"`
// }
// // NewSchedule -- create a new maintenance schedule
// func NewSchedule() *Schedule{
// 	return &Schedule{
// 		Windows: make([]*Window,0),
// 	}
// }

// FIXME: Borrowed this from flight-director, worry about implementing this later 12/6/2017
// func (d *Data) String() string {
// 	return d.User + " : " + d.Private
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
// robot robots.Robot,
func newBrain(data *Data) *Brain {
	// TODO: We need to pass in some sort of struct that represents commandline data
	// githubAPIURL, err := parseGithubAPI(config.GithubAPI)
	// if err != nil {
	// 	return nil, err
	// }

	brain := Brain{
		data:     &Data{},
		autoSave: false,
	}

	return &brain
}

// FIXME: Disabled for now 12/6/2017
// Value Receiver Method
// func (b *brain) set(newvalue string) {
// 	// b.

// 	// Grades := make(map[string]float32)

// 	// Grades["Timmy"] = 42
// 	// NOTE: If you don't make this into a pointer receiver, then setting value below will not persist other calls!
// 	c.top_speed_kmh = 500
// 	// Returns top speed of car in km/h
// 	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbixmax)
// }

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
