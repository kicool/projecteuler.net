package main

import (
	"fmt"
	//	"os"
	"bytes"
	"io/ioutil"
	"sort"
	"strings"
)

const (
	path = "./names.txt"
)

func readFile(path string) (strs []string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	strs = strings.Split(bytes.NewBuffer(file).String(), ",")
	return
}

func calcValue(strs []string) (sum int64) {
	sum = 0
	for i, v := range strs {
		sum += calcItem(strings.Trim(v, "\" ")) * int64(i+1)
	}
	return
}

func calcItem(str string) (sum int64) {
	sum = 0
	for _, v := range str {
		sum += int64(v - 'A' + 1)
	}
	return
}

func main() {
	strs := readFile(path)
	sort.Strings(strs)
	fmt.Println(calcValue(strs))

}
