package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_openLock(t *testing.T) {

	dead := []string{}
	v := 0

	dead = []string{"0201", "0101", "0102", "1212", "2002"}
	v = openLock(dead, "0202")
	assert.Equal(t, 6, v)

	dead = []string{"8888"}
	v = openLock(dead, "0009")
	assert.Equal(t, 1, v)

	dead = []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	v = openLock(dead, "8888")
	assert.Equal(t, -1, v)

}
