package main

import (
	"flag"
)

func parseArgs() (string, string, []string) {
	// 引数をパースするのにflagを使う
	logDir := flag.String(
		"logdir", "", "Log output directory (default=stderr)")
	flag.Parse()
	return *logDir, flag.Arg(0), flag.Args()[1:]
}
