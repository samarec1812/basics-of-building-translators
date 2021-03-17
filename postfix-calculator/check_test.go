package main

import (
	"./check"
	"./proc"
	"testing"
)

type testCheckPair struct {
	input   string
	correct bool
}

var testsCheck = []testCheckPair{
	{"-2+3", true},
	{"a+b", true},
	{"", false},
	{"aaaaaa", false},
	{"a", true},
	{"(2+3)-(g+n)", false},
	{"a\\b", true},
	{"(f-g)", true},
	{"((g+h)-b", false},
	{"g+(a-b)\\c+t", true},
}

func TestCorrectString(t *testing.T) {
	for _, pair := range testsCheck {
		value := check.CorrectString(pair.input)
		if value != pair.correct {
			t.Error("For", pair.input,
				"expected", pair.correct,
				"got", value)
		}
	}
}


type testPair struct {
	input    string
	expected string
}

var tests = []testPair{
	{"   -2 +  3 ", "-2+3"},
	{" a +b           ", "a+b"},
	{"", ""},
	{" aaaaaa", "aaaaaa"},
	{"   ", ""},
	{"(2+3  )- (g + n)", "(2+3)-(g+n)"},
}

func TestProcessing(t *testing.T) {
	for _, pair := range tests {
		value := proc.Processing(pair.input)
		if value != pair.expected {
			t.Error("For", pair.input,
				"expected", pair.expected,
				"got", value)
		}
	}
}
/*
-2 + 3
-2*a
*/