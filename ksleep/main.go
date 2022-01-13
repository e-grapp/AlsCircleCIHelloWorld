package main

import (
	"strconv"
	"time"

	"bitbucket.org/ai69/amoy"
	flag "github.com/spf13/pflag"
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

	// parse basic duration
	var duration time.Duration
	rawDur := args[0]
	if dur, err1 := time.ParseDuration(rawDur); err1 == nil {
		duration = dur
		log.Debugw("got basic duration", "duration", duration)
	} else if sec, err2 := strconv.ParseFloat(rawDur, 64); err2 == nil {
		duration = time.Duration(sec * float64(time.Second))
		log.Debugw("got float duration", "duration", duration)
	} else {
		log.Fatalw("invalid duration", zap.Error(err1), zap.Error(err2))
	}

	// basic mode w/o randomization
	if len(args) == 1 {
		sleep(duration)
		return
	}

	// parse random rate
	rawRate := args[1]
	rate, err := strconv.ParseFloat(rawRate, 64)
	if err != nil {
		log.Fatalw("invalid rate", zap.Error(err))
	} else if rate < 0 {
		log.Fatalw("invalid rate which should be larger than 0")
	}
	edgeDuration := time.Duration(float64(duration) * rate)

	// get random duration
	var (
		minDur, maxDur time.Duration
	)
	if duration < edgeDuration {
		minDur = duration
		maxDur = edgeDuration
	} else {
		minDur = edgeDuration
		maxDur = duration
	}
	log.Debugw("range for random duration", "min_duration", minDur, "max_duration", maxDur)
	randDur := amoy.RandomDuration(minDur, maxDur)

	// go to sleep
	sleep(randDur)
}
