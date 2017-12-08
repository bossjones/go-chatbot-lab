package main

import (
	"testing"
)

func TestFullVersion(t *testing.T) {
	version := FullVersion()

	expected := Version + VersionPrerelease + " (" + GitCommit + ")"

	if version != expected {
		t.Fatalf("invalid version returned: %s", version)
	}
}
