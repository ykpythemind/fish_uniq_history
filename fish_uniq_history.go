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
	list := makeUniqedHistory(file)
	fmt.Fprint(os.Stdout, strings.Join(list, "\n"))
}

func makeUniqedHistory(reader io.Reader) []string {
	commandList := read(reader)
	reverse(commandList)
	uniqedList := uniq(commandList)
	return uniqedList
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

func read(reader io.Reader) (historyList []string) {
	sc := bufio.NewScanner(reader)

	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			break
		}

		bytes := sc.Bytes()
		if string(bytes[:1]) != "-" {
			continue
		}

		commandStr := bytes[7:] // "- cmd : hogehoge"
		historyList = append(historyList, string(commandStr))
	}
	return
}
