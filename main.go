package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	const defaultLine = 2
	const defaultFile = "read.txt"
	var line int
	var file string
	flag.IntVar(&line, "line", defaultLine, "reading line")
	flag.StringVar(&file, "file", defaultFile, "reading file")
	flag.Parse()
	ans, err := readLines(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(ans[line])
}

func readLines(filename string) ([]string, error) {
	ans := make([]string, 10)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ans, fmt.Errorf(filename + " can't be opened")
	}
	ans = strings.Split(string(data), "\n")

	return ans, err
}
