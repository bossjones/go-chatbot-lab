package version_test

import (
	. "github.com/bossjones/go-chatbot-lab/shared/version"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Version", func() {

	Describe("Default Version variables", func() {
		It("should be <Unknown>", func() {
			Expect(VersionPrerelease).To(Equal(""))
			Expect(Version).To(Equal("<Unknown>"))
			Expect(BuildDate).To(Equal("<Unknown>"))
			Expect(GitCommit).To(Equal("<Unknown>"))
			// Expect(GoVersion).To(BeNumerically(">", "go1.9.2"))
		})
	})
})
