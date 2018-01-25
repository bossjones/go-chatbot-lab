package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

	logUtils "github.com/bossjones/go-chatbot-lab/shared/log-utils"
)

// Config -
type Config struct {
	ConfigFile string `json:"configFile"`
	// "name": "chatbot",
	Name string `json:"name"`
	// "commands": ["how are you:fine","where are you:here"],
	Commands []string `json:"commands"`
	// "db-path": "./chatbot.db",
	DBPath string `json:"db-path"`
	// "adaptor-path": "/opt/adaptors/",
	AdaptorPath string `json:"adaptor-path"`
	// "brain-type": "in-memory",
	BrainType string `json:"brain-type"`
	// "alias": "bot",
	Alias string `json:"alias"`
	// "log-level": "info",
	LogLevel string `json:"log-level"`
	// "log-location": "stdout",
	LogLocation string `json:"log-location"`
	// "log-app-name": "chatbot",
	LogAppName string `json:"log-app-name"`
	// "host": "127.0.0.1",
	Host string `json:"host"`
	// "port": "2001",
	Port string `json:"port"`
	// "kvStoreUsername": "user",
	KVUsername string `json:"kvStoreUsername"`
	// "kvStorePassword": "pass",
	KVPassword string `json:"kvStorePassword"`
	// "kvStoreServerAddress": "localhost:4001",
	KVServer string `json:"kvStoreServerAddress"`
	// "kv-ttl": "10",
	KVttl string `json:"kv-ttl"`
	// "isContainer": false,
	IsContainer bool `json:"isContainer"`
	// "ssl-cert-location": "/etc/chatbot",
	SSLCertLocation string `json:"ssl-cert-location"`
	// "log-level": "debugging level"
	Debug bool `json:"logLevel"`
}

var chatLog = logUtils.NewLogger()

// NewConfig -
func NewConfig(file string) (Config, error) {
	c := Config{}
	if file != "" {
		err := c.LoadFromConfigFile(file)
		return c, err
	}

	c.LoadFromEnv()
	return c, nil

}

// LoadFromEnv -
func (c *Config) LoadFromEnv() {
	setValueFromEnv(&c.Name, "CHATBOT_NAME")
	setSliceValueFromEnv(&c.Commands, "CHATBOT_COMMANDS")
	setValueFromEnv(&c.DBPath, "CHATBOT_DB_PATH")
	setValueFromEnv(&c.AdaptorPath, "CHATBOT_ADAPTOR_PATH")
	setValueFromEnv(&c.BrainType, "CHATBOT_BRAIN_TYPE")
	setValueFromEnv(&c.Alias, "CHATBOT_ALIAS")
	setValueFromEnv(&c.LogLevel, "CHATBOT_LOG_LEVEL")
	setValueFromEnv(&c.LogLocation, "CHATBOT_LOG_LOCATION")
	setValueFromEnv(&c.LogAppName, "CHATBOT_LOG_APP_NAME")
	setValueFromEnv(&c.Host, "CHATBOT_HOST")
	setValueFromEnv(&c.Port, "CHATBOT_PORT")
	setValueFromEnv(&c.KVUsername, "CHATBOT_KV_STORE_USERNAME")
	setValueFromEnv(&c.KVPassword, "CHATBOT_KV_STORE_PASSWORD")
	setValueFromEnv(&c.KVServer, "CHATBOT_KV_STORE_SERVER_ADDRESS")
	setValueFromEnv(&c.KVttl, "CHATBOT_KV_TTL")
	setBoolValueFromEnv(&c.IsContainer, "CHATBOT_IS_CONTAINER")
	setValueFromEnv(&c.SSLCertLocation, "CHATBOT_SSL_CERT_LOCATION")
	setBoolValueFromEnv(&c.Debug, "CHATBOT_DEBUG")
}

// LoadFromConfigFile -
func (c *Config) LoadFromConfigFile(configFile string) error {

	c.ConfigFile = configFile
	jsonSrc, err := ioutil.ReadFile(c.ConfigFile)
	if err != nil {
		chatLog.Warn("Error reading config.", "error", err)
		return err
	}
	err = json.Unmarshal(jsonSrc, &c)
	if err != nil {
		chatLog.Warn("Error parsing config.", "error", err)

		return err
	}
	return nil
}

// String() is a custom method that returns the Config without DockerPassword
func (c *Config) String() string {
	var buffer bytes.Buffer

	v := reflect.ValueOf(c).Elem()
	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Name
		val := v.Field(i).String()

		buffer.WriteString(key + ": ")
		if strings.Contains(strings.ToLower(key), "password") {
			buffer.WriteString("******" + "\n")
		} else if strings.ToLower(key) == "commands" {
			for _, i := range c.Commands {
				buffer.WriteString(i + " ")
			}
			buffer.WriteString("\n")
		} else {
			buffer.WriteString(val + "\n")
		}
	}
	return buffer.String()
}

//CommandListIsManuallySet returns true if a lidt f commands was passed in to config
// this is used to tell Capcom not to check for dynamic command list from Flight Director.
func (c *Config) CommandListIsManuallySet() bool {
	tempConfig, _ := NewConfig(c.ConfigFile)
	return (len(tempConfig.Commands) > 0)
}

func setValueFromEnv(field *string, envVar string) {
	env := os.Getenv(envVar)
	if len(env) > 0 {
		*field = env
	}
}

func setSliceValueFromEnv(field *[]string, envVar string) {
	env := os.Getenv(envVar)
	if len(env) > 0 {
		apps := make([]string, 0)
		err := json.Unmarshal([]byte(env), &apps)
		if err != nil {
			chatLog.Error("Error parsing slice in config.", "variable", envVar, "value", env)
		}
		*field = apps
	}
}

// func setIntValueFromEnv(field *int, envVar string) {
// 	env := os.Getenv(envVar)
// 	if len(env) > 0 {
// 		var err error
// 		*field, err = strconv.Atoi(env)
// 		if err != nil {
// 			chatLog.Error("Invalid environment variable", "var", envVar, "value", env)
// 		}
// 	}
// }

// //Validate checks config values for obvious problems.
// func (c *Config) Validate() error {
// 	var e error
// 	if _, err := strconv.Atoi(c.APIServerPort); err != nil {
// 		e = errors.New(APIServerPortErr + c.APIServerPort)
// 	}
// 	if !strings.Contains("sqlite3,mysql", c.DBEngine) {
// 		e = errors.New(DBEngineErr + c.DBEngine)
// 	}
// 	if c.DBPath == "" {
// 		e = errors.New(DBPathErr + c.DBPath)
// 	}
// 	//mysql needs usernmae password and database
// 	if c.DBEngine != "sqlite3" && c.DBUsername == "" {
// 		e = errors.New(DBUsernameErr + c.DBUsername)
// 	}
// 	if c.DBEngine != "sqlite3" && c.DBPassword == "" {
// 		e = errors.New(DBPasswordErr + c.DBPassword)
// 	}
// 	if c.DBEngine != "sqlite3" && c.DBDatabase == "" {
// 		e = errors.New(DBDatabaseErr + c.DBDatabase)
// 	}
// 	if c.MarathonMaster == "" {
// 		e = errors.New(MarathonMasterErr + c.MarathonMaster)
// 	}
// 	if c.MesosMaster == "" {
// 		e = errors.New(MesosMasterErr + c.MesosMaster)
// 	}
// 	if c.MesosAuthenticationPrincipal != "" && c.MesosAuthenticationSecretFile == "" {
// 		e = errors.New("Cannot have empty Mesos authentication file if principal is specified")
// 	}
// 	if c.MarathonAuthenticationPrincipal != "" && c.MarathonAuthenticationSecretFile == "" {
// 		e = errors.New("Cannot have empty Marathon authentication file if principal is specified")
// 	}
// 	if c.MarathonAuthenticationPrincipal != "" && c.MarathonUser != "" {
// 		e = errors.New("Cannot specify both MarathonUser and MarathonPrincipal")
// 	}
// 	if c.KVServer == "" {
// 		e = errors.New(KVServerErr + c.KVServer)
// 	}
// 	rangeVals := strings.Split(c.ProxyPortRange, ":")
// 	if len(rangeVals) != 2 {
// 		e = errors.New(ProxyPortRangeFormatErr + c.ProxyPortRange)
// 	} else {
// 		if _, err := strconv.Atoi(rangeVals[0]); err != nil {
// 			e = errors.New(ProxyPortRangeMinErr + c.ProxyPortRange)
// 		}
// 		if _, err := strconv.Atoi(rangeVals[1]); err != nil {
// 			e = errors.New(ProxyPortRangeMaxErr + c.ProxyPortRange)
// 		}
// 	}
// 	if !strings.Contains("json-file journald", c.AppLogDriver) {
// 		e = errors.New(LogErr + c.AppLogDriver)
// 	}
// 	if c.AquaEnabled && c.AquaEndpoint == "" {
// 		e = errors.New(AquaErr)
// 	}
// 	for _, param := range c.AllowedDockerRunParams {
// 		if strings.ContainsAny(param, "[]\"") {
// 			e = errors.New(
// 				"FD_ALLOWED_DOCKER_RUN_PARAMS should be a comma-separated list of param names; no brackets or quotes. E.g.:\n" +
// 					"    label,l,read-only,work-dir,w,network,net\n")
// 		}
// 	}
// 	return e
// }

func setBoolValueFromEnv(field *bool, envVar string) {
	env := os.Getenv(envVar)
	if len(env) > 0 {
		var err error
		*field, err = strconv.ParseBool(env)
		if err != nil {
			chatLog.Error("Invalid environment variable", "var", envVar, "value", env)
		}
	}
}
