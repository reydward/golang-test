package level3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterUnique(t *testing.T) {
	input := []Developer{
		{"Elliot"},
		{"Elliot"},
	}

	expected := []string{
		"Elliot",
	}

	assert.Equal(t, expected, FilterUnique(input))
}