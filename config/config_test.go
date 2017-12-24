package config_test

import (
	"os"
	"reflect"

	log "github.com/behance/go-logging/log"
	. "github.com/bossjones/go-chatbot-lab/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	var (
		configFile string
		c          *Config
	)

	BeforeEach(func() {
		log.SetLevel(log.PanicLevel)
		configFile = "./sampleconfig.json"
		c = &Config{}
	})

	Describe("Reading sample config file", func() {
		It("should read the file", func() {

			err := c.LoadFromConfigFile(configFile)
			log.Error(err)

			Expect(err).NotTo(HaveOccurred())
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
	Describe("Reading from ENV", func() {
		It("should read a var", func() {
			os.Setenv("CHATBOT_HOST", "a-server-name")
			c, err := NewConfig("")

			Expect(err).NotTo(HaveOccurred())
			Expect(c.Host).To(Equal("a-server-name"))
		})
		It("should read bool vars", func() {
			os.Setenv("CHATBOT_IS_CONTAINER", "false")
			c, err := NewConfig("")

			Expect(err).NotTo(HaveOccurred())
			Expect(c.IsContainer).To(Equal(false))
		})
		// It("should read proxy template resource", func() {
		// 	os.Setenv("CHATBOT_PROXY_TEMPLATE_RESOURCE", "/my/path/to.template")
		// 	c, err := NewConfig("")

		// 	Expect(err).NotTo(HaveOccurred())
		// 	Expect(c.ProxyTemplateResource).To(Equal("/my/path/to.template"))
		// })
		It("should read a slice var", func() {
			os.Setenv("CHATBOT_COMMANDS", `["app1:img1","app1:img2"]`)
			c, err := NewConfig("")

			Expect(err).NotTo(HaveOccurred())
			Expect(len(c.Commands)).To(Equal(2))
			Expect(c.Commands[1]).To(Equal("app1:img2"))
		})
		It("should  fail to read a bad slice var", func() {
			os.Setenv("CHATBOT_COMMANDS", `"app1:img1","app1:img2"`) //missing []
			c, _ := NewConfig("")
			Expect(len(c.Commands)).To(Equal(0))
		})
		// It("should read marathon variables", func() {
		// 	os.Setenv("CHATBOT_MARATHON_EVENT_HANDLER", "true")
		// 	os.Setenv("CHATBOT_MARATHON_MASTER", "marathon-master")
		// 	os.Setenv("CHATBOT_MARATHON_MASTER_PROTOCOL", "marathon-master-protocol")
		// 	os.Setenv("CHATBOT_MARATHON_USER", "marathon-user")
		// 	os.Setenv("CHATBOT_MARATHON_PASSWORD", "marathon-password")
		// 	os.Setenv("CHATBOT_MARATHON_AUTH_PRINCIPAL", "marathon-auth-principal")
		// 	os.Setenv("CHATBOT_MARATHON_SECRET_FILE", "marathon-secret-file")
		// 	c, err := NewConfig("")

		// 	Expect(err).NotTo(HaveOccurred())
		// 	Expect(c.MarathonEventHandler).To(BeTrue())
		// 	Expect(c.MarathonMaster).To(Equal("marathon-master"))
		// 	Expect(c.MarathonMasterProtocol).To(Equal("marathon-master-protocol"))
		// 	Expect(c.MarathonUser).To(Equal("marathon-user"))
		// 	Expect(c.MarathonPassword).To(Equal("marathon-password"))
		// 	Expect(c.MarathonAuthPrincipal).To(Equal("marathon-auth-principal"))
		// 	Expect(c.MarathonSecretFile).To(Equal("marathon-secret-file"))
		// })
	})
})
