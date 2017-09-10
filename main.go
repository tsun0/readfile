package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

/*
const BUFSIZE = 1024 // 読み込みバッファのサイズ

func main() {
	file, err := os.Open(`read.txt`)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	buf := make([]byte, BUFSIZE)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			// Readエラー処理
			break
		}

		fmt.Print(string(buf[:n]))
	}
}
*/

func main() {
	const defaultLine = 2
	var line int
	flag.IntVar(&line, "line", defaultLine, "line to use")
	flag.Parse()
	ans, err := readLines("read.txt")
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
