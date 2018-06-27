package main

import (
	"testing"
)

var testFile = "test_history"

var expectCommandList = [12]string{"brew install ghq", "docker --version", "ghq get https://github.com/ykpythemind/dotfiles.git", "brew doctor", "sh build.sh ", "chsh", "sudo vim /etc/shells", "man curl", "sh init.sh ", "vim .zshrc", "vim .zshenv "}

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
