package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"

	"github.com/dickeyy/adventofcode/2015/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("./input.txt")
	utils.Output(day4(i, p))
}

func day4(key string, part int) int {
	num := 0
	z := "00000"
	if part == 2 {
		z = "000000"
	}

	for {
		num++
		hash := calcMD5(key + strconv.Itoa(num))
		if strings.HasPrefix(hash, z) {
			return num
		}
	}
}

// helper func to simplify md5 calculation
func calcMD5(s string) string {
	h := md5.New() // thank god go has a built in md5 hash function
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
