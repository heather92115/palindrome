package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected string
	}{

		{"asd", "dsa"},
		{"bananas", "sananab"},
	}
	for _, test := range tests {
		actual := Reverse(test.param)
		if actual != test.expected {
			t.Errorf("Expected Reverse(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestFindPalindrome(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected string
	}{

		{"dsa", ""},
		{"bananas", "anana"},
		{"good\n", "oo"},
		{"nothingtofind", ""},
		{"aviddiva", "aviddiva"},
		{"aviddivakl", "aviddiva"},
		{"test", ""},
		{"Aeratepetarea", "aeratepetarea"},
		{"wwanna", "anna"},
		{"uptonewera", "ewe"},
	}
	for _, test := range tests {
		actual := FindPalindrome(test.param, 0)
		if actual != test.expected {
			t.Errorf("Expected Palindrome (%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}