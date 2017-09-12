package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// -hオプション用文言
	flag.Usage = func() {
		fmt.Fprintf(
			os.Stderr,
			`
Usage of %s:
  %s [OPTIONS] ARGS...
Options
`,
			os.Args[0],
			os.Args[0],
		)
		flag.PrintDefaults()
	}
	const defaultLine = 0
	var line int
	var file string
	flag.IntVar(&line, "line", defaultLine, "reading line")
	flag.StringVar(&file, "file", "", "reading file")
	flag.Parse()
	ans, err := readLines(file)
	if err != nil {
		log.Println(file + ": file can't be opened")
	}
	fmt.Println(ans[line])
}

func readLines(filename string) ([]string, error) {
	ans := make([]string, 10)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ans, err
	}
	ans = strings.Split(string(data), "\n")

	return ans, err
}
