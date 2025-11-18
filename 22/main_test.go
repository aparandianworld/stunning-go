package main

import (
	"strings"
	"testing"
)

var (
	benchSink []string
)

func BenchmarkTokenize(b *testing.B) {
	text := "Do not communicate by sharing memory, share memory by communicating"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		benchSink = Tokenize(text)
	}
}

func TestTokenize(t *testing.T) {
	text := "Do not communicate by sharing memory, share memory by communicating"
	expected := []string{"do", "not", "communicate", "by", "sharing", "memory", "share", "memory", "by", "communicating"}

	tokens := Tokenize(text)

	if !strings.EqualFold(strings.Join(tokens, " "), strings.Join(expected, " ")) {
		t.Errorf("Expected %v, got %v", expected, tokens)
	}
	
}
