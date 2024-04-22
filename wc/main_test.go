package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	input := bytes.NewBufferString("word1 word2 word3 word4\n")
	want := 4

	AssertCount(t, input, want, "")
}

func TestCountLines(t *testing.T) {
	input := bytes.NewBufferString("word1\nword2\nword3")
	want := 3

	AssertCount(t, input, want, "l")
}

func TestCountBytes(t *testing.T) {
	input := bytes.NewBufferString("word1")
	exp := 5

	AssertCount(t, input, exp, "b")
}

func AssertCount(t testing.TB, input *bytes.Buffer, want int, flagType string) {
	t.Helper()

	var got int

	switch flagType {
	case "l":
		got = count(input, true, false)
	case "b":
		got = count(input, false, true)
	default:
		got = count(input, false, false)
	}

	if got != want {
		t.Errorf("Expected %d, got %d instead.\n", got, want)
	}
}
