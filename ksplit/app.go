package main

import (
	"os"

	"bitbucket.org/ai69/amoy"
	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
)

var (
	lineLimit  int
	nameFormat string
	startIndex int

	log      *zap.SugaredLogger
	logLevel string
	version  bool
)

func init() {
	flag.IntVarP(&lineLimit, "limit", "l", 0, "limit of lines for each file")
	flag.StringVarP(&nameFormat, "format", "f", `split_%03d.txt`, "format of file name")
	flag.IntVarP(&startIndex, "start", "s", 0, "start index of file name")

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
