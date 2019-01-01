package main

import (
	"log"
	"os"
	"testing"
)

var testFile = "testdata/test_history"

func TestRead(t *testing.T) {
	file, err := os.Open(testFile)
	if err != nil {
		log.Fatalf("[Error] %v\n", err)
	}
	defer file.Close()
	list := read(file)

	expectCommandList := []string{"cmd1", "cmd2", "cmd3", "cmd4", "cmd5", "cmd2", "cmd6", "cmd7", "cmd2", "cmd3", "cmd2"}

	for i, actual := range list {
		expect := expectCommandList[i]
		if actual != expect {
			t.Errorf("failed: expected %s, but got %s", expect, actual)
		}
	}
}

func BenchmarkRead(b *testing.B) {
	file, err := os.Open(testFile)
	if err != nil {
		log.Fatalf("[Error] %v\n", err)
	}
	defer file.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		read(file)
	}
	b.StopTimer()
}

func TestMakeUniqedHistory(t *testing.T) {
	file, err := os.Open(testFile)
	if err != nil {
		log.Fatalf("[Error] %v\n", err)
	}
	defer file.Close()

	var expect = []string{"cmd2", "cmd3", "cmd7", "cmd6", "cmd5", "cmd4", "cmd1"}
	list := makeUniqedHistory(file)
	for i := 0; i < len(list); i++ {
		if expect[i] != list[i] {
			t.Errorf("failed: expect %s, but got %s", expect[i], list[i])
		}
	}
}

func TestUniq(t *testing.T) {
	dupList := []string{"a", "aaa", "a", "piyo", "hoge", "hoge", "a", "piyo2", "piyo", "fuga"}
	expectList := []string{"a", "aaa", "piyo", "hoge", "piyo2", "fuga"}
	uniqedList := uniq(dupList)
	for i := 0; i < len(expectList); i++ {
		if expectList[i] != uniqedList[i] {
			t.Errorf("failed: expect %s, but got %s", expectList[i], uniqedList[i])
		}
	}
}

func TestReverse(t *testing.T) {
	list := []string{"1", "2", "4", "a", "5", "3"}
	expectList := []string{"3", "5", "a", "4", "2", "1"}
	reverse(list)
	for i := 0; i < len(expectList); i++ {
		if expectList[i] != list[i] {
			t.Errorf("failed: expect %s, but got %s", expectList[i], list[i])
		}
	}
}
