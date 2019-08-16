package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeTurnedLines(t *testing.T) {
	testCases := []struct {
		in       string
		in2      int
		expected []string
	}{
		{"Golang", 6, []string{"Golang"}},
		{"Golang", 7, []string{"Golang "}},
		{"Golang", 5, []string{"Golan", "g    "}},
	}

	for _, c := range testCases {
		t.Run(fmt.Sprintf("Input(str:%s, maxLen:%d)", c.in, c.in2), func(t *testing.T) {
			// Act
			act := MakeTurnedLines(c.in, c.in2)
			// Assert
			assert.Equal(t, c.expected, act)
		})
	}
}
