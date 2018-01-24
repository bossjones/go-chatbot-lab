package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

	log "github.com/behance/go-logging/log"
)

// Config -
type Config struct {
	// ********************************************
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
	// ********************************************
}

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
}

// LoadFromConfigFile -
func (c *Config) LoadFromConfigFile(configFile string) error {

	c.ConfigFile = configFile
	jsonSrc, err := ioutil.ReadFile(c.ConfigFile)
	if err != nil {
		log.Warn("Error reading config.", "error", err)
		return err
	}
	err = json.Unmarshal(jsonSrc, &c)
	if err != nil {
		log.Warn("Error parsing config.", "error", err)

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
			log.Error("Error parsing slice in config.", "variable", envVar, "value", env)
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
// 			log.Error("Invalid environment variable", "var", envVar, "value", env)
// 		}
// 	}
// }

func setBoolValueFromEnv(field *bool, envVar string) {
	env := os.Getenv(envVar)
	if len(env) > 0 {
		var err error
		*field, err = strconv.ParseBool(env)
		if err != nil {
			log.Error("Invalid environment variable", "var", envVar, "value", env)
		}
	}
}
