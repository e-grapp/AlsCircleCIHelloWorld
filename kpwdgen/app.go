package main

import (
	"os"

	"bitbucket.org/ai69/amoy"
	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
)

var (
	log      *zap.SugaredLogger
	logLevel string
	version  bool

	count      int
	length     int
	numDigits  int
	numSymbols int
	allLower   bool
)

func init() {
	flag.BoolVarP(&version, "version", "v", false, "show version")
	flag.StringVar(&logLevel, "log", "info", "set log level")
	flag.IntVarP(&count, "count", "c", 1, "number of passwords to generate")
	flag.IntVarP(&length, "length", "l", 16, "password length")
	flag.IntVarP(&numDigits, "digits", "d", 2, "number of digits")
	flag.IntVarP(&numSymbols, "symbols", "s", 1, "number of symbols")
	flag.BoolVar(&allLower, "lower", false, "all lowercase")
	flag.Parse()
	if version {
		os.Stderr.WriteString(buildInfo())
		os.Exit(0)
	}

	lg := amoy.NewPersistentLogger("", true)
	lg.SetLogLevel(logLevel)
	log = lg.LoggerSugared()
}
