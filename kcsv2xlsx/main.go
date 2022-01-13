package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"bitbucket.org/ai69/amoy"
	"github.com/1set/gut/ystring"
	flag "github.com/spf13/pflag"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
)

func sleep(d time.Duration) {
	log.Infow("will sleep for a while", "duration", d)
	time.Sleep(d)
	log.Infow("now wake up")
}

func main() {
	args := flag.Args()
	if len(args) <= 0 {
		log.Fatal("no arguments")
	}

	// open csv
	path := args[0]
	cf, err := os.Open(path)
	if err != nil {
		log.Fatalw("failed to open csv file", "path", path, zap.Error(err))
	}
	defer cf.Close()

	// read csv
	cr := csv.NewReader(cf)
	lines, err := cr.ReadAll()
	if err != nil {
		log.Fatalw("failed to read csv file", "path", path, zap.Error(err))
	}
	log.Infow("read csv file", "path", path, "records", len(lines))

	if len(lines) <= 0 {
		return
	}

	// create xlsx
	if ystring.IsBlank(outputFile) {
		outputFile = ystring.NotBlankOrDefault(amoy.SubstrBeforeLast(path, "."), "output") + ".xlsx"
	}

	// create a new sheet
	ef := excelize.NewFile()
	sheetName := "Sheet1"

	width := len(lines[0])
	maxCol, _ := excelize.ColumnNumberToName(width)
	_ = ef.SetColWidth(sheetName, "A", maxCol, colWidth)

	// set values for each line
	for idx, line := range lines {
		cell := fmt.Sprintf("A%d", idx+1)
		err := ef.SetSheetRow(sheetName, cell, &line)
		if err != nil {
			log.Fatalw("fail to write row", "cell", cell, "line", line, zap.Error(err))
		}
	}

	// save xlsx
	if err := ef.SaveAs(outputFile); err != nil {
		log.Fatalw("fail to save xlsx", "path", outputFile, zap.Error(err))
	}
	log.Infow("saved xlsx file", "path", outputFile, "column_width", colWidth)
}
