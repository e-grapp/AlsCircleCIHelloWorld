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
)

func init() {
	flag.BoolVarP(&version, "version", "v", false, "show version")
	flag.StringVar(&logLevel, "log", "info", "set log level")
	flag.Parse()
	if version {
		os.Stderr.WriteString(buildInfo())
		os.Exit(0)
	}

	lg := amoy.NewPersistentLogger("", true)
	lg.SetLogLevel(logLevel)
	log = lg.LoggerSugared()
}
