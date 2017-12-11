package lib

import (
	"github.com/twinj/uuid"
)

// import "encoding/json"

// https://siongui.github.io/2016/01/30/go-pretty-print-variable/

// PrettyPrint - function that takes an interface
// func PrettyPrint(v interface{}) {
// 	b, _ := json.MarshalIndent(v, "", "  ")
// 	println(string(b))
// }

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
