package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubArray(t *testing.T) {

	assert.Equal(t, 1, minInsertions("(()))"))

	assert.Equal(t, 0, minInsertions("())"))

	assert.Equal(t, 3, minInsertions("))())"))
}
