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

func TestData(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Data Suite")
}

var _ = Describe("Brain", func() {
	Describe("Creating a Data object using NewData", func() {

		It("Should return a Data object with property User and Private of type {<str>:<str>}", func() {
			d := *NewData()

			Expect(d.User).To(BeAssignableToTypeOf(make(map[string]string)))
			Expect(d.Private).To(BeAssignableToTypeOf(make(map[string]string)))
		})

		It("Should return a Data object of type Data", func() {
			d := *NewData()

			Expect(&d).To(BeAssignableToTypeOf(new(Data)))
		})

	})

	Describe("Creating a Brain using NewBrain", func() {
		It("Should return a Brain object", func() {
			d := *NewData()
			b := *NewBrain(&d)

			Expect(&b).To(BeAssignableToTypeOf(new(Brain)))
		})
	})

})
