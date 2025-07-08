package cli

import (
	"advent/utils/debug"
	"fmt"
	"slices"
	"strings"
)

// NOTE that mutable values must be pointers
type ArgsList struct {
	Choices []string
	Val *string
	DefaultVal string
}

func NewArgsList(choices []string, defaultVal string) *ArgsList {
	if choices == nil {
		debug.DebugPrintln("must provide choices")
		return nil
	}
	if !slices.Contains(choices, defaultVal) {
		debug.DebugPrintln("default %s not in choices %s", defaultVal, choices)
		return nil
	}
	return &ArgsList{Choices: choices, Val: &defaultVal, DefaultVal: defaultVal}
}

func (argsList ArgsList) String() string {
	if argsList.Val == nil {
		return ""
	}
	return *argsList.Val
}

func (argsList ArgsList) Set(val string) error {
	if !slices.Contains(argsList.Choices, val) {
		return fmt.Errorf("val %s not in choices %s", val, argsList.Choices)
	}

	*argsList.Val = val
	return nil
}

func (argsList ArgsList) Usage() string {
	return fmt.Sprintf("--mode=[%s]", strings.Join(argsList.Choices, "|"))
}
