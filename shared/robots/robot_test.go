package robots_test

import (
	"fmt"
	"testing"

	. "github.com/bossjones/go-chatbot-lab/shared/brains"
	. "github.com/bossjones/go-chatbot-lab/shared/robots"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// INFO: https://stackoverflow.com/questions/6478962/what-does-the-dot-or-period-in-a-go-import-statement-do

// If an explicit period (.) appears instead of a name, all the package's exported identifiers will be declared in the current file's file block and can be accessed without a qualifier.

// Assume we have compiled a package containing the package clause package math, which exports function Sin, and installed the compiled package in the file identified by "lib/math". This table illustrates how Sin may be accessed in files that import the package after the various types of import declaration.

func TestData(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Data Suite")
}

var _ = Describe("Robot", func() {
	Describe("Creating a Robot object using NewRobot", func() {

		It("Should return a Robot object using default values", func() {
			// adapterpath *string, adapter *string, httpd *bool, name *string, alias *string

			adapterpath_str := string("shell")
			adapter_str := string("shell")
			httpd_bool := bool(false)
			name_str := string("Scarlett")
			alias_str := string("Scarlett")
			r := *NewRobot(&adapterpath_str, &adapter_str, &httpd_bool, &name_str, &alias_str)

			// AdapterPath: adapterpath,
			// Adapter:     adapter,
			// Httpd:       httpd,
			// Name:        name,
			// Alias:       alias,
			// RobotBrain:  rbrain,

			// ********************************************************************
			// NOTE: Needed a reminder of what my current values are
			fmt.Println("*****AIYO*******")
			fmt.Println(adapterpath_str)
			fmt.Println(&adapterpath_str)
			fmt.Println(r.AdapterPath)
			fmt.Println(*r.AdapterPath)
			fmt.Println("*****AIYO*******")

			// *****AIYO*******
			// shell
			// 0xc42004d440
			// 0xc42004d440
			// shell
			// *****AIYO*******
			// ********************************************************************

			Expect(*r.AdapterPath).To(BeEquivalentTo("shell"))
			Expect(*r.Adapter).To(BeEquivalentTo("shell"))
			Expect(*r.Httpd).To(BeFalse())
			Expect(*r.Name).To(BeEquivalentTo("Scarlett"))
			Expect(*r.Alias).To(BeEquivalentTo("Scarlett"))
			Expect(r.RobotBrain).To(BeAssignableToTypeOf(new(Brain)))

		})

	})

})
