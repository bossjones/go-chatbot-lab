package brains_test

// INFO: https://stackoverflow.com/questions/6478962/what-does-the-dot-or-period-in-a-go-import-statement-do

// If an explicit period (.) appears instead of a name, all the package's exported identifiers will be declared in the current file's file block and can be accessed without a qualifier.

// Assume we have compiled a package containing the package clause package math, which exports function Sin, and installed the compiled package in the file identified by "lib/math". This table illustrates how Sin may be accessed in files that import the package after the various types of import declaration.

// "testing"

import (
	"testing"

	. "github.com/bossjones/go-chatbot-lab/shared/brains"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestData(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Data Suite")
}

var _ = Describe("Data", func() {
	var (
		data Data
	)

	BeforeEach(func() {
		data = *NewData()
	})
	Describe("NewData", func() {
		It("Is initialized properly", func() {
			Expect(data.User).To(BeAssignableToTypeOf(make(map[string]string)))
			Expect(data.Private).To(BeAssignableToTypeOf(make(map[string]string)))
		})
	})

})
