package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashing(t *testing.T) {
	testTable := []struct {
		name     string
		msg      string
		arg      string
		equal    bool
		expected string
	}{
		{
			name:     "OK",
			msg:      "Should be equal",
			arg:      "google.com",
			equal:    true,
			expected: "huPVAz7b6R",
		},
		{
			name:     "Another one OK",
			msg:      "Should be equal",
			arg:      "github.com",
			equal:    true,
			expected: "uoJDg5WT5K",
		},
		{
			name:     "Compare results of different inputs",
			msg:      "Should not be equal",
			arg:      "google.com",
			equal:    false,
			expected: hashing("yandex.ru"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			res := hashing(test.arg)
			if test.equal != true {
				assert.NotEqual(t, test.expected, res, test.msg)
			} else {
				assert.Equal(t, test.expected, res, test.msg)
			}
		})
	}
}
