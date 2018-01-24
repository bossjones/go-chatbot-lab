package lib

import (
	"bytes"
	"os/exec"
	"time"

	"github.com/behance/go-common/log"
	"github.com/twinj/uuid"
)

// import "encoding/json"

// https://siongui.github.io/2016/01/30/go-pretty-print-variable/

// PrettyPrint - function that takes an interface
// func PrettyPrint(v interface{}) {
// 	b, _ := json.MarshalIndent(v, "", "  ")
// 	println(string(b))
// }

// Coalesce takes string arguments and returns the first one that is not empty
// This is useful for coalscing information from multiple sources
func Coalesce(args ...string) string {
	for i := 0; i < len(args); i++ {
		if args[i] != "" {
			return args[i]
		}
	}
	return ""
}

// CoalesceDuration takes duration arguments and returns the first one that is not empty
// This is useful for coalescing information from multiple sources
func CoalesceDuration(args ...time.Duration) time.Duration {
	for i := 0; i < len(args); i++ {
		if args[i] != 0 {
			return args[i]
		}
	}
	return 0
}

// CoalesceBool takes bool arguments and returns the first one that is not false
// This is useful for coalescing information from multiple sources
func CoalesceBool(args ...bool) bool {
	for i := 0; i < len(args); i++ {
		if args[i] == true {
			return args[i]
		}
	}
	return false
}

// CoalesceInt takes int arguments and returns the first one that is not zero
// This is useful for coalescing information from multiple sources
func CoalesceInt(args ...int) int {
	for i := 0; i < len(args); i++ {
		if args[i] > 0 {
			return args[i]
		}
	}
	return 0
}

// Iff is a go implemetation of ternary operations (condition ? iftrue : if false)
func Iff(condition bool, ifTrue interface{}, ifFalse interface{}) interface{} {
	if condition {
		return ifTrue
	}
	return ifFalse
}

// In - checks whether the string is in the array
func In(val string, targ []string) bool {
	for _, cur := range targ {
		if cur == val {
			return true
		}
	}
	return false
}

// CreateUUID returns a string version of a V1 uuid
func CreateUUID() string {
	u1 := uuid.NewV1()
	return u1.String()
}

// ExecuteShellScript execs a specified script file and returns the output in a string
func ExecuteShellScript(scriptFile string) (string, error) {
	cmd := exec.Command("bash", "-c", scriptFile)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		log.Error("Script execution error", "action", "execute shell script", "command", cmd, "error", err, "error_details", stderr.String())
		return stderr.String(), err

	}
	log.Info("Executed: "+scriptFile, "action", "execute shell script", "details", out.String())

	return out.String(), nil
}
