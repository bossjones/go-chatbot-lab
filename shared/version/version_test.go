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

	Describe("Func FullVersion", func() {
		It("returns a properly formatted string", func() {
			full_version := FullVersion()
			expected_full_version := "<Unknown> (<Unknown>)"

			Expect(full_version).To(Equal(expected_full_version))
		})
	})

	Describe("Func DetailedVersionInfo", func() {
		It("returns a properly formatted string", func() {
			detailed_version_info := DetailedVersionInfo()
			expected_detailed_version_info := "Go-Chatbot-Lab <Unknown>; buildDate=<Unknown>; sha=<Unknown>"

			Expect(detailed_version_info).To(Equal(expected_detailed_version_info))
		})
	})

})
