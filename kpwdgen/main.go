package main

import (
	"fmt"

	"github.com/sethvargo/go-password/password"
	"go.uber.org/zap"
)

func main() {
	var result []string
	for i := 0; i < count; i++ {
		res, err := password.Generate(length, numDigits, numSymbols, allLower, false)
		if err != nil {
			log.Warn("fail to generate password", "index", i, zap.Error(err))
		}
		result = append(result, res)
	}
	log.Debugw("generated passwords", "count", count, "length", length, "numDigits", numDigits, "numSymbols", numSymbols, "allLower", allLower, "result", result)

	for _, p := range result {
		fmt.Println(p)
	}
}
