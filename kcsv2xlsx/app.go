package main

import (
	"os"

	"bitbucket.org/ai69/amoy"
	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
)

var (
	outputFile string
	colWidth   float64
	log        *zap.SugaredLogger
	logLevel   string
	version    bool
)

func init() {
	flag.StringVarP(&outputFile, "output", "o", "", "output file name")
	flag.Float64VarP(&colWidth, "col-width", "w", 60, "set column width")

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
