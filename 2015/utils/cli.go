package utils

import (
	"flag"
	"fmt"
)

var part int

const (
	colorReset = "\033[0m"
	darkGray   = "\033[90m"
	pink       = "\033[35m"
)

func init() {
	flag.IntVar(&part, "p", 1, "part 1 or 2")
}

func ParseFlags() {
	flag.Parse()
	fmt.Printf("%sRunning part %d...%s\n", darkGray, part, colorReset)
}

func GetPart() int {
	return part
}

func Output(data interface{}) {
	println("Output: ", pink, fmt.Sprintf("%v%s", data, darkGray)) // ends in dark gray so that if you use the time command it looks nicer
	// eventually ill add other stuff here like copying to clipboard and stuff
}
