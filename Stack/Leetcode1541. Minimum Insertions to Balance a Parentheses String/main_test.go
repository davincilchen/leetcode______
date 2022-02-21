package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubArray(t *testing.T) {

	assert.Equal(t, 1, minInsertions("(()))"))

	assert.Equal(t, 0, minInsertions("())"))

	assert.Equal(t, 3, minInsertions("))())("))

	assert.Equal(t, 12, minInsertions("(((((("))

	assert.Equal(t, 3, minInsertions("))))))"))
	assert.Equal(t, 5, minInsertions(")))))))"))

	assert.Equal(t, 3, minInsertions("()())))()"))

	assert.Equal(t, 1, minInsertions("(()))"))
	assert.Equal(t, 4, minInsertions("(()))(()))()())))"))

	assert.Equal(t, 16, minInsertions("))(()()))()))))))()())()(())()))))()())(()())))()("))

}
