package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	file, err := os.Open(historyFilePath())
	if err != nil {
		log.Fatalf("[Error] %v\n", err)
	}
	defer file.Close()
	fmt.Fprint(os.Stdout, makeUniqedHistory(file))
}

func makeUniqedHistory(reader io.Reader) string {
	commandList := read(reader)
	reverse(commandList)
	uniqedList := uniq(commandList)
	return strings.Join(uniqedList, "\n")
}

func historyFilePath() string {
	return filepath.Join(os.Getenv("HOME"), ".local", "share", "fish", "fish_history")
}

func uniq(list []string) []string {
	uniqedList := make([]string, len(list))
	m := make(map[string]bool)
	var i int
	for _, v := range list {
		_, ok := m[v]
		if ok {
			continue
		}
		m[v] = true
		uniqedList[i] = v
		i++
	}
	return uniqedList[:i]
}

func reverse(list []string) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}

var prefix = []byte("-")

func read(reader io.Reader) []string {
	sc := bufio.NewScanner(reader)
	var historyList []string

	for sc.Scan() {
		if err := sc.Err(); err != nil {
			break
		}

		bytes := sc.Bytes()
		if bytes[0] != prefix[0] {
			continue
		}

		// https://github.com/fish-shell/fish-shell/blob/63e70d601d449b0b1448f63f58e2db25576d1822/src/history.cpp#L610
		commandStr := bytes[7:] // "- cmd : hogehoge"
		historyList = append(historyList, string(commandStr))
	}
	return historyList
}
