package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubArray(t *testing.T) {

	assert.Equal(t, 1, costToBalance("(()))"))

	assert.Equal(t, 0, costToBalance("())"))

	assert.Equal(t, 3, costToBalance("))())("))

}
