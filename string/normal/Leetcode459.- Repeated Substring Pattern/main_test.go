package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_repeatedSubstringPattern(t *testing.T) {

	assert.Equal(t, true, repeatedSubstringPattern("abab"))

	assert.Equal(t, false, repeatedSubstringPattern("aba"))
	assert.Equal(t, true, repeatedSubstringPattern("abcabcabcabc"))
}
