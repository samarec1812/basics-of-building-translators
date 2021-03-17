package proc

import "strings"


// This func erase all space symbol
func Processing(expr string) string {
	s := strings.ReplaceAll(expr, " ", "")
	return s
}
