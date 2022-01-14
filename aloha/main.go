package main

import (
	"fmt"
	"time"
)

func main() {
	// pwd, _ := os.Getwd()
	// host, _ := os.Hostname()
	// log.Infow("Hello Go World", "hostname", host, "work_dir", pwd)
	fmt.Printf("🌋: Hello World!\n⏰: %s\n", time.Now().Format("2006-01-02T15:04:05-0700"))

	fmt.Println(`	┌─┐┬  ┌─┐┬ ┬┌─┐
	├─┤│  │ │├─┤├─┤
	┴ ┴┴─┘└─┘┴ ┴┴ ┴`)
}
