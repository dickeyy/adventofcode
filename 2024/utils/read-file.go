package util

import (
	"os"
	"strings"
)

/*
Example Usage of ReadFile:
data := utils.ReadFile("./data.txt")
*/

func ReadFile(path string) string {
	// open the file, panic if error
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close() // close the file when done
	// read the data from the file, panic if error
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	// trim whitespace from the data and return
	t := strings.TrimSpace(string(data))
	return t
}
