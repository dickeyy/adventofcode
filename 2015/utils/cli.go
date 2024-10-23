package utils

import (
	"flag"
	"fmt"
	"time"
)

var (
	part  int
	timer bool
)

const (
	colorReset = "\033[0m"
	darkGray   = "\033[90m"
	pink       = "\033[35m"
)

func init() {
	flag.IntVar(&part, "p", 1, "part 1 or 2")
	flag.BoolVar(&timer, "t", false, "time the solution")
}

func ParseFlags() {
	flag.Parse()
	fmt.Printf("%sRunning part %d...%s\n", darkGray, part, colorReset)
	if timer {
		startTime = time.Now()
	}
}

func GetPart() int {
	return part
}

func Output(data interface{}) {
	var elapsed time.Duration
	if timer {
		elapsed = time.Since(startTime)
	}

	if timer {
		fmt.Printf("%sOutput: %s%v %s(in %v)%s\n",
			darkGray, pink, data, darkGray, elapsed, colorReset)
	} else {
		fmt.Printf("%sOutput: %s%v%s\n",
			darkGray, pink, data, colorReset)
	}
}

// StartTimer starts timing if timer flag is enabled
var startTime time.Time

func StartTimer() {
	if timer {
		startTime = time.Now()
	}
}
