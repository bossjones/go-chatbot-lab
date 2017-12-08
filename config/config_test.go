package config_test

// import (
// 	"testing"
// 	log "github.com/bossjones/go-chatbot-lab/log"
// 	. "github.com/bossjones/go-chatbot-lab/config"
// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// )

// var _ = Describe("Server/Engine", func() {
// 	var (
// 		engine Engine
// 		cfg    Config
// 		kv     KV
// 	)

// 	BeforeEach(func() {
// 		log.SetLevel(log.PanicLevel)
// 		cfg = Config{}
// 		engine = *NewEngine(cfg)

// 		kv.Store = kvwrapper.NewKVWrapper([]string{""}, kvwrapper.KVFaker{})
// 		kv.UniqueMachineID = "aaaaaaaaaa"
// 		testHelper.Init()
// 		testHelper.SaveProxyToKV(kv)
// 	})

// 	Describe("NewEngine", func() {
// 		It("Is initialized properly", func() {
// 			Expect(engine.Active).To(BeFalse())
// 			Expect(engine.QuitWork).To(BeAssignableToTypeOf(make(chan bool)))
// 		})
// 	})

// 	Context("Engine is already running", func() {
// 		BeforeEach(func() {
// 			engine.Active = true
// 		})

// 		Describe("Engine#Start", func() {
// 			It("Will not start again", func() {
// 				err := engine.Start()
// 				Expect(err).To(Equal(ErrEngineAlreadyStarted))
// 			})
// 		})
// 	})
// })
