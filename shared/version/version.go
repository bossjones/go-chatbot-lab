package version

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"
)

// // GitCommit - The git commit that was compiled. This will be filled in by the compiler.
// var GitCommit string

// // Version - The main version number that is being run at the moment.
// const Version = "0.1.0"

// VersionPrerelease - A pre-release marker for the version. If this is "" (empty string)
// then it means that it is a final release. Otherwise, this is a pre-release
// such as "dev" (in development)
var VersionPrerelease = ""

// Version of this chatbot binary
var Version = "<Unknown>"

// BuildDate is the build date of this chatbot binary
var BuildDate = "<Unknown>"

// GitCommit is the git SHA of this chatbot binary
var GitCommit = "<Unknown>"

// GoVersion is the version of Golang used to build flight-director
var GoVersion = runtime.Version()

// FullVersion -
func FullVersion() string {
	return Version + VersionPrerelease + " (" + GitCommit + ")"
}

// DetailedVersionInfo returns a string with a bunch of info about Flight
// Director, meant for putting in a `Server` response header, a `User-Agent`
// request header, etc.
func DetailedVersionInfo() string {
	return fmt.Sprintf(
		"Flight Director %s; buildDate=%s; sha=%s",
		Version, BuildDate, GitCommit,
	)
}

// ConvertToNumeric converts a version string (. separated) into a numeric value for easier > and < comparisons
func ConvertToNumeric(version string) float64 {
	parts := strings.Split(version, ".")
	var number float64
	for i, s := range parts {
		part, _ := strconv.Atoi(s)
		number += math.Pow(1000, float64(len(parts)-(i+1))) * float64(part)
	}

	return number
}
