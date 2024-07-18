package main

import (
	"strings"
	"testing"
)

func detend(in string) string {
	lines := strings.Split(strings.TrimSpace(in), "\n")
	for idx := range lines {
		lines[idx] = strings.TrimSpace(lines[idx])
	}
	return strings.Join(lines, "\n")
}

func testTextToCols(t *testing.T, in, expected string) {
	actual := textToCols(in, 4)

	if actual != expected {
		t.Fatalf("\nExected:\n%s\nActual:\n%s", expected, actual)
	}
}

func TestTextToCols1(t *testing.T) {
	in := detend(`1
                  2
                  3
                  4
                  5
                  6
                  7
                  8
                  9
                  10
                  11
                  12`)

	expected := detend(`1    4    7    10
                        2    5    8    11
                        3    6    9    12`)

	testTextToCols(t, in, expected)
}
