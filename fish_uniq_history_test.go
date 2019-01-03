package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var (
	testFile     = "testdata/history"
	testLongFile = "testdata/long_history"
)

func TestRead(t *testing.T) {
	file, err := os.Open(testFile)
	if err != nil {
		log.Fatalf("[Error] %v\n", err)
	}
	defer file.Close()

	h := newHistory(file)
	list := h.read()

	expectCommandList := []string{"cmd1", "cmd2", "cmd3", "cmd4", "cmd5", "cmd2", "cmd6", "cmd7", "cmd2", "cmd3", "cmd2"}

	for i, actual := range list {
		expect := expectCommandList[i]
		if actual != expect {
			t.Errorf("failed: expected %s, but got %s", expect, actual)
		}
	}
}

func BenchmarkReadLong(b *testing.B) {
	file, err := os.Open(testLongFile)
	if err != nil {
		log.Fatalf("[Error] %v\n", err)
	}
	defer file.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := newHistory(file)
		h.read()
	}
	b.StopTimer()
}

func BenchmarkRead(b *testing.B) {
	file, err := os.Open(testFile)
	if err != nil {
		log.Fatalf("[Error] %v\n", err)
	}
	defer file.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := newHistory(file)
		h.read()
	}
	b.StopTimer()
}

func BenchmarkOutput(b *testing.B) {
	file, err := os.Open(testFile)
	if err != nil {
		log.Fatalf("[Error] %v\n", err)
	}
	defer file.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := newHistory(file)
		err = h.output(ioutil.Discard)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
}

func TestOutput(t *testing.T) {
	file, err := os.Open(testFile)
	if err != nil {
		log.Fatalf("[Error] %v\n", err)
	}
	defer file.Close()

	h := newHistory(file)
	b := new(bytes.Buffer) // A Buffer needs no initialization.
	err = h.output(b)
	if err != nil {
		t.Error(err)
	}

	expect := "cmd2\ncmd3\ncmd7\ncmd6\ncmd5\ncmd4\ncmd1"
	if b.String() != expect {
		t.Errorf("failed: expect: %s, but got %s", expect, b.String())
	}
}

func TestMakeUniqList(t *testing.T) {
	file, err := os.Open(testFile)
	if err != nil {
		log.Fatalf("[Error] %v\n", err)
	}
	defer file.Close()

	h := newHistory(file)
	list := h.makeUniqedList()

	var expect = []string{"cmd2", "cmd3", "cmd7", "cmd6", "cmd5", "cmd4", "cmd1"}
	for i, v := range expect {
		if v != list[i] {
			t.Errorf("failed: expect %s, but got %s", v, list[i])
		}
	}
}

func TestReverseUniq(t *testing.T) {
	dupList := []string{"a", "aaa", "a", "piyo", "hoge", "hoge", "a", "piyo2", "piyo", "fuga"}
	expectList := []string{"fuga", "piyo", "piyo2", "a", "hoge", "aaa"}
	newList := reverseUniq(dupList)
	for i := 0; i < len(expectList); i++ {
		if expectList[i] != newList[i] {
			t.Errorf("failed: expect %s, but got %s", expectList[i], newList[i])
		}
	}
}
