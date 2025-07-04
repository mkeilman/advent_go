package cli

import (
	"advent/utils/debug"
	"fmt"
	"slices"
)

type ArgsList struct {
	Choices []string
	Val string
}

func NewArgsList(choices []string, initVal string) *ArgsList {
	if choices == nil {
		debug.DebugPrintln("must provide choices")
		return nil
	}
	if !slices.Contains(choices, initVal) {
		debug.DebugPrintln("val %s not in choices %s", initVal, choices)
		return nil
	}
	return &ArgsList{Choices: choices, Val: initVal}
}

func (argsList ArgsList) String() string {
	return argsList.Val
}

func (argsList ArgsList) Set(val string) error {
	if !slices.Contains(argsList.Choices, val) {
		return fmt.Errorf("val %s not in choices %s", val, argsList.Choices)
	}

	argsList.Val = val
	debug.DebugPrintln("SET %s", argsList.Val)
	return nil
}
