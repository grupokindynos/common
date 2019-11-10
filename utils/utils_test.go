package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	array := []string{"a", "b"}
	contains := Contains(array, "a")
	assert.Equal(t, true, contains)
}

func TestContainsError(t *testing.T) {
	array := []string{"a", "b"}
	contains := Contains(array, "e")
	assert.Equal(t, false, contains)
}