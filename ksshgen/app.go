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

	algoType   string
	bitsNumber int
	fileName   string
	override   bool
	passphrase string
	title      string
)

func init() {
	flag.StringVarP(&algoType, "algo", "a", "rsa", "algorithm type: rsa, dsa, ecdsa and ed25519")
	flag.IntVarP(&bitsNumber, "bits", "b", 0, "number of bits in the key to create")
	flag.StringVarP(&fileName, "file", "f", "", "file name to save the key")
	flag.BoolVarP(&override, "override", "o", false, "override the file if it exists")
	flag.StringVarP(&passphrase, "passphrase", "p", "", "set a passphrase for the key")
	flag.StringVarP(&title, "title", "t", "", "set a title for the key")

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
