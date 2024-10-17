package utils

import (
	"flag"
	"fmt"
)

var part int

func init() {
	flag.IntVar(&part, "p", 1, "part 1 or 2")
}

func ParseFlags() {
	flag.Parse()
	fmt.Println("Running part", part)
}

func GetPart() int {
	return part
}
