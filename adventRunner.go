package main

import (
	"advent/utils/debug"
	"advent/utils/cli"
	"advent/utils/collections"

	"errors"
	"flag"
	"fmt"
	"os"
	"slices"

	"golang.org/x/tools/go/packages"
)

const (
	ERR_NUM_ARGS = iota + 90
	ERR_INVALID_YEAR
)

var MODES = []string{"test", "file", "all"}

func packageByName(name string) (*packages.Package, error) {

	// must do recursive search - Load itself does NOT fail if it looks for packages that do not exist
	pk, e := packages.Load(&packages.Config{Mode: packages.NeedName}, "./...")
	if e != nil {
		return nil, e
	}

	// now filter the results
	pk = collections.Filter(pk, func (p *packages.Package) bool { return p.Name == name })
	if len(pk) == 0 {
		return nil, errors.New("Package missing")
	}

	return pk[0], nil
}

func main() {

	m := cli.NewArgsList(MODES, "all")

	usage := func() {
		fmt.Fprintf(os.Stderr, "Usage: go run advent %s <year> <day>\n", m.Usage())
	}

	flag.Var(*m, "mode", "run mode")
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		usage()
		os.Exit(ERR_NUM_ARGS)
	}

	year, _ := args[0], args[1]

	_, err := packageByName(fmt.Sprintf("year%s", year))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot load package for year %s: %s\n", args[0], err)
		os.Exit(ERR_INVALID_YEAR)
	}

	//pk, e := packages.Load(&packages.Config{Mode: packages.NeedName}, fmt.Sprintf("filename=day%s.go", day))

	//d := New(2024, 11)
	//var d Runner

	if slices.Contains([]string{"test", "all"}, *m.Val) {
		debug.DebugPrintln("TEST:")
		//d.run_from_test_input()
	}

	if slices.Contains([]string{"file", "all"}, *m.Val) {
		debug.DebugPrintln("FILE:")
		//d.run_from_file()
	}

	//debug.DebugPrintln("A: %s F: %s", os.Args[1:], flag.Args())
}
