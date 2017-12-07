package prompts

// Borrowed from https://github.com/dixonwille/wmenu/blob/master/clearScreen.go

import (
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("/usr/bin/clear") // nolint: errcheck
		cmd.Stdout = os.Stdout
		cmd.Run() // nolint: errcheck
	}
	clear["darwin"] = func() {
		cmd := exec.Command("/usr/bin/clear") // nolint: errcheck
		cmd.Stdout = os.Stdout
		cmd.Run() // nolint: errcheck
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") // nolint: errcheck
		cmd.Stdout = os.Stdout
		cmd.Run() // nolint: errcheck
	}
}

//Clear simply clears the command line interface (os.Stdout only).
func Clear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	}
}
