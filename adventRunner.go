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
	//"strconv"

	"golang.org/x/tools/go/packages"
)

const (
	ERR_NUM_ARGS = iota + 90
	ERR_INVALID_YEAR
	ERR_INVALID_DAY
)

var MODES = []string{"test", "file", "all"}
var modes = cli.NewArgsList(MODES, "all")

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

func exit(code int) {
	usage()
	os.Exit(code)
}

func main() {

	//m := cli.NewArgsList(MODES, "all")

	//usage := func() {
	//	fmt.Fprintf(os.Stderr, "Usage: go run advent %s <year> <day>\n", m.Usage())
	//}

	flag.Var(*modes, "mode", "run mode")
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		exit(ERR_NUM_ARGS)
	}

	/*
	year, err := strconv.ParseInt(args[0], 10, 0)
	if err != nil {
		exit(ERR_INVALID_YEAR)
	}
	day, err := strconv.ParseInt(args[1], 10, 0)
	if err != nil {
		exit(ERR_INVALID_DAY)
	}
	*/

	/*
	pkg, err := packageByName(fmt.Sprintf("year%d", year))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot load package for year %s: %s\n", args[0], err)
		os.Exit(ERR_INVALID_YEAR)
	}
	*/

	//pk, e := packages.Load(&packages.Config{Mode: packages.NeedName}, fmt.Sprintf("filename=day%s.go", day))

	//d := New(int(year), int(day))
	//d.RunFromFile()
	//var d Runner

	if slices.Contains([]string{"test", "all"}, *modes.Val) {
		debug.DebugPrintln("TEST:")
		//d.RunFromTestInput()
	}

	if slices.Contains([]string{"file", "all"}, *modes.Val) {
		debug.DebugPrintln("FILE:")
		//d.run_from_file()
	}

	//debug.DebugPrintln("A: %s F: %s", os.Args[1:], flag.Args())
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: go run advent %s <year> <day>\n", modes.Usage())
}