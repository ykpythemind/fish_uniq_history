package main

import (
	"testing"
)

var testFile = "test_history"

var expectCommandList = [12]string{"brew install ghq", "docker --version", "ghq get https://github.com/ykpythemind/dotfiles.git", "brew doctor", "sh build.sh ", "docker --version", "sudo vim /etc/shells", "man curl", "sh init.sh ", "vim .zshrc", "vim .zshenv "}

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
