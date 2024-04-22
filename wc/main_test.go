package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	input := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4
	res := count(input, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	input := bytes.NewBufferString("word1\nword2\nword3")

	exp := 3
	res := count(input, true)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
