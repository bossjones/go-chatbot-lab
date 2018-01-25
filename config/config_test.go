package config_test

import (
	"os"
	"reflect"

	. "github.com/bossjones/go-chatbot-lab/config"
	logUtils "github.com/bossjones/go-chatbot-lab/shared/log-utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var chatLog = logUtils.NewLogger()

func setenv(key, val string) {
	err := os.Setenv(key, val)
	if err != nil {
		chatLog.Error("Setenv Error", "err", err)
	}
}

func unsetenv(key string) {
	err := os.Unsetenv(key)
	if err != nil {
		chatLog.Error("Unsetenv Error", "err", err)
	}
}

var _ = Describe("Config", func() {
	var (
		configFile string
		c          *Config
	)

	BeforeEach(func() {
		configFile = "./sampleconfig.json"
		c = &Config{}
	})

	Describe("Reading sample config file", func() {
		It("should read the file", func() {
			err := c.LoadFromConfigFile(configFile)
			Expect(err).ToNot(HaveOccurred())
		})
		It("it should parse json correctly", func() {
			var noValue string
			_ = c.LoadFromConfigFile(configFile)
			v := reflect.ValueOf(c).Elem()
			for i := 0; i < v.NumField(); i++ {
				key := v.Type().Field(i).Name
				val := v.Field(i).String()

				if val == "" {
					noValue = noValue + ", " + key
				}
			}
			Expect(noValue).To(BeEmpty())
		})

	})

	Describe("LoadFromEnv", func() {
		BeforeEach(func() {
			setenv("CHATBOT_HOST", "a-server-name")
			setenv("CHATBOT_COMMANDS", `["app1:img1","app1:img2"]`)
			setenv("CHATBOT_IS_CONTAINER", "false")
		})

		It("reads ENV vals", func() {
			c.LoadFromEnv()
			Expect(c.Host).To(Equal("a-server-name"))
			Expect(c.IsContainer).To(Equal(false))
			Expect(len(c.Commands)).To(Equal(2))
			Expect(c.Commands[0]).To(Equal("app1:img1"))
			Expect(c.Commands[1]).To(Equal("app1:img2"))
		})

		AfterEach(func() {
			unsetenv("CHATBOT_HOST")
			unsetenv("CHATBOT_COMMANDS")
			unsetenv("CHATBOT_IS_CONTAINER")
		})
	})

	Describe("String", func() {
		It("Returns the Config without DockerPassword", func() {
			config := Config{
				Name:      "scarlett",
				BrainType: "redis",
			}

			safeConfig := config.String()

			Expect(safeConfig).To(ContainSubstring(`BrainType: redis`))
			Expect(safeConfig).To(ContainSubstring(`Name: scarlett`))
		})
	})
})
