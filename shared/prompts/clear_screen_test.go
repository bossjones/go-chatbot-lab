package prompts

import "testing"

// TestClear -
func TestClear(t *testing.T) {
	Clear()
}

// TestClearLinux -
func TestClearLinux(t *testing.T) {
	clearOs("linux")
}

// TestClearDarwin -
func TestClearDarwin(t *testing.T) {
	clearOs("darwin")
}

// TestClearWindows -
func TestClearWindows(t *testing.T) {
	clearOs("windows")
}

// clearOs -
func clearOs(os string) {
	value, ok := clear[os]
	if ok {
		value()
	}
}
