package lib

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
