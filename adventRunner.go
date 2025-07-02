package main

import (
	"advent/utils/debug"
	"flag"
	"slices"
)

func main() {

	m := ""
	flag.StringVar(&m, "mode", "all", "run mode")

	//p.add_argument("year", type=int, help="year to run")
	//p.add_argument("day", type=int, help="day to run")
	//p.add_argument(
	//    "--mode",
	//    type=str,
	//    help="run mode",
	//    choices=["test", "file", "all"],
	//    default="all",
	//    dest="mode",
	//)
	//a, u = p.parse_known_args()
	//d = day.Build(a.year, a.day, u)
	//d = day.Build(2024, 11)

	//modes := []string{"test", "file", "all"}

	//flag.wordPtr(&d.mode)

	flag.Parse()
	debug.DebugPrintln("MODE: %s", m)

	//if !slices.Contains(m, modes) {

	//}

	if slices.Contains([]string{"test", "all"}, m) {
		debug.DebugPrintln("TEST:")
		//d.run_from_test_input()
	}

	if slices.Contains([]string{"file", "all"}, m) {
		debug.DebugPrintln("FILE:")
		//d.run_from_file()
	}
}
