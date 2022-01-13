package main

import (
	"fmt"

	"bitbucket.org/ai69/amoy"
	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
)

func main() {
	args := flag.Args()
	if len(args) <= 0 {
		log.Fatal("no input files")
	}

	var (
		err     error
		lines   []string
		cntLine int
		cntFile = startIndex
	)
	checkFileWrite := func(force bool) {
		if len(lines) == 0 {
			return
		}
		if force || (cntLine >= lineLimit && lineLimit > 0) {
			name := fmt.Sprintf(nameFormat, cntFile)
			if err = amoy.WriteFileLines(name, lines); err != nil {
				log.Errorw("failed to write file", "file", name, zap.Error(err))
			}
			log.Infow("saved split file", "file", name, "count", cntFile, "lines", cntLine)
			cntFile++
			cntLine = 0
			lines = []string{}
		}
	}

	for idx, f := range args {
		if err = amoy.ReadFileByLine(f, func(line string) (err error) {
			lines = append(lines, line)
			cntLine++
			checkFileWrite(false)
			return nil
		}); err != nil {
			log.Errorw("failed to read file", "index", idx, "file", f, zap.Error(err))
		}
	}
	checkFileWrite(true)
}
