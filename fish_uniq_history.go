package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	ls := Read(historyFile())
	uniqedList := Uniq(&ls)
	// fmt.Println(strings.Join(ls, "\n"))
	Out(&uniqedList)
}

func historyFile() string {
	return filepath.Join(os.Getenv("HOME"), ".local", "share", "fish", "fish_history")
}

func Out(historyList *[]string) {
	fmt.Fprint(os.Stdout, strings.Join(*historyList, "\n"))
}

func Uniq(historyList *[]string) (uniqedList []string) {
	mapList := make(map[string]bool)
	for _, v := range *historyList {
		_, ok := mapList[v]
		if ok {
			continue
		}
		mapList[v] = true
		uniqedList = append(uniqedList, v)
	}
	for i, j := 0, len(uniqedList)-1; i < j; i, j = i+1, j-1 {
		uniqedList[i], uniqedList[j] = uniqedList[j], uniqedList[i]
	}
	return
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
