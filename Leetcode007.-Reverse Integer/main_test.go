package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {

	v := 0

	v = reverse(123)
	assert.Equal(t, 321, v)

	v = reverse(-123)
	assert.Equal(t, -321, v)

	v = reverse(120)
	assert.Equal(t, 21, v)

	v = reverse(1534236469)
	assert.Equal(t, 0, v)

	v = reverse(1563847412)
	assert.Equal(t, 0, v)

}
