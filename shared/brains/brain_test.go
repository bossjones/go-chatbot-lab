package brains_test

// INFO: https://stackoverflow.com/questions/6478962/what-does-the-dot-or-period-in-a-go-import-statement-do

// If an explicit period (.) appears instead of a name, all the package's exported identifiers will be declared in the current file's file block and can be accessed without a qualifier.

// Assume we have compiled a package containing the package clause package math, which exports function Sin, and installed the compiled package in the file identified by "lib/math". This table illustrates how Sin may be accessed in files that import the package after the various types of import declaration.

import (
	"testing"
	. "github.com/bossjones/go-chatbot-lab/shared/brains"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


var _ = Describe("Data", func() {
	var (
		d Data
	)

	BeforeEach(func() {
		// log.SetLevel(log.PanicLevel)
		d = NewData()
	})
	// Describe("Data Object", func() {
	// 	It("User property is a empty dictonary of strings", func() {
	// 		Expect(d).To(BeEmpty())
	// 	})
	// })
	Describe("Data String representation", func() {
		It("User string representation is BLAH", func() {
			Expect(d.User).To(BeEmpty())
		})
	})

})

// type Data struct {
// 	// User is a struct field that is of type map {"":""}
// 	User          map[string]string

// 	// Private is a struct field that is of type map {"":""}
// 	Private       map[string]string
// }
