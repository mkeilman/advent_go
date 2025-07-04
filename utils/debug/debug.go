package debug

import (
	"fmt"
)

//Debugging utilities

// A wrapper to avoid "naked prints"
//
// Args:
//
//	txtFormat (string): something to print, properly formatted
//	end (string): string to append
//	args (any): printf args
func DebugPrint(txtFormat string, end string, args ...any) {
	fmt.Print(fmt.Sprintf(txtFormat, args...) + fmt.Sprintf("%s", end))
}

// Print with a newline
//
// Args:
//
//	txtFormat (string): something to print, properly formatted
//	args (any): printf args
func DebugPrintln(txtFormat string, args ...any) {
	DebugPrint(txtFormat, "\n", args...)
}

// Conditional print
//
// Args:
//
//	txtFormat (string): something to print, properly formatted
//	condition (bool): if true, print; otherwise do nothing
func DebugIf(txtFormat string, end string, condition bool, args ...any) {
	if condition {
		DebugPrint(txtFormat, end, args...)
	}
}
