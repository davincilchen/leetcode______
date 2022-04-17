package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubArray(t *testing.T) {

	assert.Equal(t, 1, minAddToMakeValid("(()))"))

	assert.Equal(t, 1, minAddToMakeValid("())"))

	assert.Equal(t, 4, minAddToMakeValid("))())("))

	assert.Equal(t, 6, minAddToMakeValid("(((((("))

	assert.Equal(t, 6, minAddToMakeValid("))))))"))
	assert.Equal(t, 7, minAddToMakeValid(")))))))"))

}

func Test_maxSubArray2(t *testing.T) {

	// assert.Equal(t, 1, minAddToMakeValid2("(()))"))

	// assert.Equal(t, 1, minAddToMakeValid2("())"))

	// assert.Equal(t, 4, minAddToMakeValid2("))())("))

	// assert.Equal(t, 6, minAddToMakeValid2("(((((("))

	// assert.Equal(t, 6, minAddToMakeValid2("))))))"))
	// assert.Equal(t, 7, minAddToMakeValid2(")))))))"))

}
