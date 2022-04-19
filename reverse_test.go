package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, tc := range testcases {
		rev, err1 := Reverse(tc.in)
		if err1 != nil {
			t.Skip("Reverse failed")
		}
		if rev != tc.want {
			t.Errorf("Reverse(%q) == %q, want %q", tc.in, rev, tc.want)
		}
	}

}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			t.Skip("Reverse failed")
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			t.Skip("Reverse(Reverse(s)) failed")
		}
		t.Logf("Number of runs: orig=%d, rev=%d, doubleRev=%d",
			utf8.RuneCountInString(orig),
			utf8.RuneCountInString(rev),
			utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
