package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
)

type history struct {
	reader io.Reader
}

func newHistory(reader io.Reader) *history {
	return &history{
		reader: reader,
	}
}

func main() {
	file, err := os.Open(historyFilePath())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	h := newHistory(file)

	err = h.output(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *history) output(w io.Writer) error {
	list := h.makeUniqedList()
	last := len(list)
	for i, s := range list {
		s := s
		if i != last-1 {
			s = s + "\n"
		}
		_, err := w.Write([]byte(s))
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *history) makeUniqedList() []string {
	list := h.read()
	return reverseUniq(list)
}

func historyFilePath() string {
	return filepath.Join(os.Getenv("HOME"), ".local", "share", "fish", "fish_history")
}

func reverseUniq(list []string) []string {
	uniqedList := make([]string, len(list))

	m := make(map[string]struct{})

	// reverse reading
	var j int
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
		_, ok := m[v]
		if ok {
			continue
		}
		m[v] = struct{}{}
		uniqedList[j] = v
		j++
	}

	return uniqedList[:j]
}

func (h *history) read() []string {
	sc := bufio.NewScanner(h.reader)
	var list []string

	for sc.Scan() {
		if err := sc.Err(); err != nil {
			break
		}

		bytes := sc.Bytes()
		if string(bytes[0]) != "-" {
			continue
		}

		// https://github.com/fish-shell/fish-shell/blob/63e70d601d449b0b1448f63f58e2db25576d1822/src/history.cpp#L610
		command := string(bytes[7:]) // "- cmd : hogehoge"
		list = append(list, command)
	}
	return list
}
