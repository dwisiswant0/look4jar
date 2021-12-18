package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/projectdiscovery/gologger"
)

func init() {
	opt = options{}

	flag.StringVar(&opt.Path, "p", "", "")
	flag.StringVar(&opt.Path, "path", "", "")

	flag.BoolVar(&opt.Verbose, "v", false, "")
	flag.BoolVar(&opt.Verbose, "verbose", false, "")

	fmt.Fprint(os.Stderr, aurora.Cyan(MSG_BANNER))
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, MSG_USAGE)
	}
	flag.Parse()

	if opt.Path == "" {
		flag.Usage()
		os.Exit(3)
	}

	gologger.DefaultLogger.SetMaxLevel(5)
	if opt.Verbose {
		gologger.DefaultLogger.SetMaxLevel(6)
	}
}
