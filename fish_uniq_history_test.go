package main

import (
	"testing"
)

var testFile = "testdata/test_history"

var expectCommandList = [12]string{"cmd1", "cmd2", "cmd3", "cmd4", "cmd5", "cmd2", "cmd6", "cmd7", "cmd2", "cmd3", "cmd2"}

func TestRead(t *testing.T) {
	list := Read(testFile)

	for i, actual := range list {
		expect := expectCommandList[i]
		if actual != expect {
			t.Errorf("failed: expected %s, but got %s", expect, actual)
		}
	}
}

func BenchmarkRead(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Read(testFile)
	}
	b.StopTimer()
}

func TestMakeList(t *testing.T) {
	var expect = [7]string{"cmd2", "cmd3", "cmd7", "cmd6", "cmd5", "cmd4", "cmd1"}
	list := MakeList(testFile)
	for i := 0; i < len(list); i++ {
		if expect[i] != list[i] {
			t.Errorf("failed: expect %s, but got %s", expect[i], list[i])
		}
	}
}

func TestUniq(t *testing.T) {
	dupList := []string{"a", "aaa", "a", "piyo", "hoge", "hoge", "a", "piyo2", "piyo", "fuga"}
	expectList := []string{"a", "aaa", "piyo", "hoge", "piyo2", "fuga"}
	uniqedList := Uniq(dupList)
	for i := 0; i < len(expectList); i++ {
		if expectList[i] != uniqedList[i] {
			t.Errorf("failed: expect %s, but got %s", expectList[i], uniqedList[i])
		}
	}
}

func TestReverse(t *testing.T) {
	list := []string{"1", "2", "4", "a", "5", "3"}
	expectList := []string{"3", "5", "a", "4", "2", "1"}
	Reverse(list)
	for i := 0; i < len(expectList); i++ {
		if expectList[i] != list[i] {
			t.Errorf("failed: expect %s, but got %s", expectList[i], list[i])
		}
	}
}
