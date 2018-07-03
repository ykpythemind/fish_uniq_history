package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	list := MakeList(historyFile())
	fmt.Fprint(os.Stdout, strings.Join(list, "\n"))
}

func MakeList(historyFile string) []string {
	commandList := Read(historyFile)
	Reverse(commandList)
	uniqedList := Uniq(commandList)
	return uniqedList
}

func historyFile() string {
	return filepath.Join(os.Getenv("HOME"), ".local", "share", "fish", "fish_history")
}

func Uniq(historyList []string) (uniqedList []string) {
	mapList := make(map[string]bool)
	for _, v := range historyList {
		_, ok := mapList[v]
		if ok {
			continue
		}
		mapList[v] = true
		uniqedList = append(uniqedList, v)
	}
	return
}

func Reverse(list []string) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}

func Read(filePath string) (historyList []string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error] %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

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
